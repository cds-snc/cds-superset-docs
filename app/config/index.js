const config = {
  port: parseInt(process.env.PORT) || 5000,
  wordpress: {
    url: process.env.WORDPRESS_URL,
    user: process.env.WORDPRESS_USER,
    password: process.env.WORDPRESS_PASSWORD,
    get authToken() {
      return Buffer.from(`${this.user}:${this.password}`).toString("base64");
    },
    menuIds: {
      en: process.env.MENU_ID_EN,
      fr: process.env.MENU_ID_FR,
    },
  },
  site: {
    names: {
      en: process.env.SITE_NAME_EN,
      fr: process.env.SITE_NAME_FR,
    },
  },
  routing: {
    pathSegmentsAllowed: parseInt(process.env.PATH_SEGMENTS_ALLOWED) || 3,
    get pathPattern() {
      return Array(this.pathSegmentsAllowed)
        .fill("/:path")
        .map((p, i) => p + (i + 1) + "?")
        .join("");
    },
  },
};

export default config;
