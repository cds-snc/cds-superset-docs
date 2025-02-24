import express, { Request, Response } from "express";
import WordPressService from "../services/wordpress";
import config from "../config";

const wordPressService = new WordPressService(config.wordpress);
const router = express.Router();

router.get(
  config.routing.pathPattern,
  async (req: Request, res: Response): Response => {
    const { lang = "en" } = req.query;
    const pathSegments: string[] = Object.values(req.params).filter(
      Boolean,
    ) as string[];
    const lastSegment: string = pathSegments[pathSegments.length - 1] || "home";

    try {
      const [page, menuItems] = await Promise.all([
        wordPressService.getPage(lastSegment, lang),
        wordPressService.getMenu(lang),
      ]);

      if (page && menuItems) {
        return res.render("page", {
          page,
          menuItems,
          isHome: lastSegment === "home",
          langSwap: lang === "en" ? "fr" : "en",
          langSwapSlug: lang === "en" ? page?.slug_fr : page?.slug_en,
          siteName: config.site.names[lang],
        });
      } else {
        console.log(`Page not found for: ${lastSegment}`);
        return res.status(404).send("Page not found");
      }
    } catch (error) {
      console.error("Route handler error:", error);
      return res.status(500).send("Internal server error");
    }
  },
);

export default router;
