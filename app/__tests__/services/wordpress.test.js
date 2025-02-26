import WordPressService from "../../services/wordpress";

global.fetch = jest.fn();

describe("services/wordpress", () => {
  const mockConfig = {
    url: "https://mock.wordpress.ca",
    authToken: "mock-token",
    menuIds: { en: "1", fr: "2" },
  };

  let service;

  beforeEach(() => {
    service = new WordPressService(mockConfig);
    jest.clearAllMocks();
  });

  describe("getPage", () => {
    it("should return the first page when available", async () => {
      fetch.mockResolvedValueOnce({
        json: async () => [{ title: "Test Page" }],
      });
      const result = await service.getPage("test-slug", "en");
      expect(fetch).toHaveBeenCalledWith(
        "https://mock.wordpress.ca/wp-json/wp/v2/pages?slug=test-slug&lang=en",
      );
      expect(result).toEqual({ title: "Test Page" });
    });

    it("should return null when no pages found", async () => {
      fetch.mockResolvedValueOnce({
        json: async () => [],
      });
      const result = await service.getPage("missing-page", "en");
      expect(result).toBeNull();
    });

    it("should throw an error if fetch fails", async () => {
      fetch.mockRejectedValueOnce(new Error("Network error"));
      await expect(service.getPage("test-slug", "en")).rejects.toThrow(
        "Network error",
      );
    });
  });

  describe("getMenu", () => {
    it("should return a structured menu", async () => {
      const mockMenuItems = [
        {
          id: 1,
          parent: 0,
          url: "https://mock.wordpress.ca/home",
          children: [],
        },
        {
          id: 2,
          parent: 1,
          url: "https://mock.wordpress.ca/sub",
          children: [],
        },
      ];
      fetch.mockResolvedValueOnce({
        json: async () => mockMenuItems,
      });
      const menu = await service.getMenu("en");
      expect(fetch).toHaveBeenCalledWith(
        "https://mock.wordpress.ca/wp-json/wp/v2/menu-items?menus=1",
        { headers: { Authorization: "Basic mock-token" } },
      );
      expect(menu.length).toBe(1);
      expect(menu[0].url).toBe("/home");
      expect(menu[0].children[0].url).toBe("/sub");
    });

    it("should throw an error if fetch fails to get menu", async () => {
      fetch.mockRejectedValueOnce(new Error("Menu fetch error"));
      await expect(service.getMenu("en")).rejects.toThrow("Menu fetch error");
    });
  });

  describe("createMenuTree", () => {
    it("should return an empty array when given no menu items", () => {
      const service = new WordPressService(mockConfig);
      const result = service["createMenuTree"]([]);
      expect(result).toEqual([]);
    });

    it("should remove base URL from menu item URLs and build a tree", () => {
      const service = new WordPressService(mockConfig);
      const menuItems = [
        {
          id: 1,
          parent: 0,
          url: "https://mock.wordpress.ca/home",
          children: [],
        },
        {
          id: 2,
          parent: 1,
          url: "https://mock.wordpress.ca/about",
          children: [],
        },
        {
          id: 3,
          parent: 0,
          url: "https://mock.wordpress.ca/contact",
          children: [],
        },
      ];

      const result = service["createMenuTree"](menuItems);

      expect(result.length).toBe(2);
      expect(result[0].url).toBe("/home");
      expect(result[0].children[0].url).toBe("/about");
      expect(result[1].url).toBe("/contact");
    });

    it("should correctly nest child menu items under the right parent", () => {
      const service = new WordPressService(mockConfig);
      const menuItems = [
        {
          id: 10,
          parent: 0,
          url: "https://mock.wordpress.ca/home",
          children: [],
        },
        {
          id: 11,
          parent: 10,
          url: "https://mock.wordpress.ca/about",
          children: [],
        },
        {
          id: 12,
          parent: 10,
          url: "https://mock.wordpress.ca/about/team",
          children: [],
        },
        {
          id: 13,
          parent: 0,
          url: "https://mock.wordpress.ca/contact",
          children: [],
        },
      ];

      const result = service["createMenuTree"](menuItems);

      expect(result.length).toBe(2);
      expect(result[0].children.length).toBe(2);
      expect(result[1].children.length).toBe(0);
    });
  });
});
