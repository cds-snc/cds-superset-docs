interface WordPressConfig {
  url: string;
  authToken?: string;
  menuIds: {
    en: string;
    fr: string;
  };
}

interface WordPressPage {
  [key: string]: string;
}

interface MenuItem {
  id: number;
  parent: number;
  url: string;
  children: MenuItem[];
}

class WordPressService {
  constructor(private config: WordPressConfig) {
    this.config = config;
  }

  async getPage(slug: string, lang: string): Promise<WordPressPage | null> {
    try {
      const response = await fetch(
        `${this.config.url}/wp-json/wp/v2/pages?slug=${slug}&lang=${lang}`,
      );
      const page: WordPressPage[] = (await response.json()) as WordPressPage[];
      return page.length ? page[0] : null;
    } catch (error: unknown) {
      console.error(
        "Error fetching page:",
        error instanceof Error ? error.message : error,
      );
      throw error;
    }
  }

  async getMenu(lang: string): Promise<MenuItem[]> {
    const menuId =
      lang === "en" ? this.config.menuIds.en : this.config.menuIds.fr;
    try {
      const response = await fetch(
        `${this.config.url}/wp-json/wp/v2/menu-items?menus=${menuId}`,
        {
          headers: {
            Authorization: `Basic ${this.config.authToken}`,
          },
        },
      );
      const menuItems: MenuItem[] = (await response.json()) as MenuItem[];
      return this.createMenuTree(menuItems);
    } catch (error: unknown) {
      console.error(
        "Error fetching menu:",
        error instanceof Error ? error.message : error,
      );
      throw error;
    }
  }

  private createMenuTree(menuItems: MenuItem[]): MenuItem[] {
    const menuTree: MenuItem[] = [];
    const menuMap: { [id: number]: MenuItem } = {};

    menuItems.forEach((item) => {
      item.url = item.url.replace(this.config.url, "");
      menuMap[item.id] = { ...item, children: [] };
    });

    menuItems.forEach((item) => {
      if (item.parent === 0) {
        menuTree.push(menuMap[item.id]);
      } else if (menuMap[item.parent]) {
        menuMap[item.parent].children.push(menuMap[item.id]);
      }
    });

    return menuTree;
  }
}

export default WordPressService;
