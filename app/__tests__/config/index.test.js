describe("config", () => {
  let config;
  const originalEnv = { ...process.env };

  beforeAll(async () => {
    process.env.WORDPRESS_URL = "https://mock.wordpress.ca";
    process.env.WORDPRESS_USER = "test-user";
    process.env.WORDPRESS_PASSWORD = "test-pass";
    process.env.MENU_ID_EN = "111";
    process.env.MENU_ID_FR = "222";
    process.env.SITE_NAME_EN = "SiteEN";
    process.env.SITE_NAME_FR = "SiteFR";
    process.env.PORT = "4000";
    process.env.PATH_SEGMENTS_ALLOWED = "4";

    const configPath = require.resolve("../../config");
    delete require.cache[configPath];
    config = (await import("../../config")).default;
  });

  afterAll(() => {
    process.env = originalEnv;
  });

  it("parses the port from environment variables", () => {
    expect(config.port).toBe(4000);
  });

  it("creates a correct authToken", () => {
    const token = Buffer.from("test-user:test-pass").toString("base64");
    expect(config.wordpress.authToken).toBe(token);
  });

  it("builds a dynamic path pattern based on PATH_SEGMENTS_ALLOWED", () => {
    expect(config.routing.pathSegmentsAllowed).toBe(4);
    expect(config.routing.pathPattern).toBe("/:path1?/:path2?/:path3?/:path4?");
  });
});
