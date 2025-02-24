// Mark these environment variables as required.
declare module "bun" {
  interface Env {
    WORDPRESS_URL: string;
    WORDPRESS_USER: string;
    WORDPRESS_PASSWORD: string;
    MENU_ID_EN: string;
    MENU_ID_FR: string;
    SITE_NAME_EN: string;
    SITE_NAME_FR: string;
    PORT: string;
    PATH_SEGMENTS_ALLOWED: string;
  }
}

const config = {
  port: parseInt(Bun.env.PORT || "5000"),
  wordpress: {
    url: Bun.env.WORDPRESS_URL,
    user: Bun.env.WORDPRESS_USER,
    password: Bun.env.WORDPRESS_PASSWORD,
    get authToken() {
      return Buffer.from(`${this.user}:${this.password}`).toString("base64");
    },
    menuIds: {
      en: Bun.env.MENU_ID_EN,
      fr: Bun.env.MENU_ID_FR,
    },
  },
  site: {
    names: {
      en: Bun.env.SITE_NAME_EN,
      fr: Bun.env.SITE_NAME_FR,
    },
  },
  routing: {
    pathSegmentsAllowed: parseInt(Bun.env.PATH_SEGMENTS_ALLOWED || "3"),
    pathPattern: "/:path1/:path2/:path3?",
  },
};

// Build the dynamic path pattern based on the number of allowed segments.
// This determines how many levels of nesting a page in WordPress can have
// and still resolve in this site.
config.routing.pathPattern = Array(config.routing.pathSegmentsAllowed)
  .fill("/:path")
  .map((p, i) => p + (i + 1) + "?")
  .join("");

export default config;
