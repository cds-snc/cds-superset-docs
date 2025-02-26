import config from "../config";

const templateHelpers = {
  dateFormat: (date) =>
    new Date(date).toLocaleDateString("en-CA", {
      year: "numeric",
      month: "numeric",
      day: "numeric",
    }),
  eq: (a, b) => a == b,
  updateMarkup: (content) => {
    if (!content) return "";

    const siteUrl = new RegExp(`${config.wordpress.url}`, "g");

    // Updates the WordPress content to match with the expected markup
    // for the Design System components.
    return content
      .replace(
        /<details class="alert alert-([^"]+)" open><summary class="h3"><h3>([^<]+)<\/h3><\/summary>(.+)<\/details>/g,
        '<section class="mt-300 mb-300"><gcds-notice type="$1" notice-title-tag="h2" notice-title="$2"><gcds-text>$3</gcds-text></gcds-notice></section>',
      )
      .replace(
        /<div class="wp-block-button"><a class="wp-block-button__link[^"]+" href="([^"]+)">([^<]+)<\/a><\/div>/g,
        '<gcds-button type="link" href="$1">$2</gcds-button>',
      )
      .replace(
        /<details class="wp-block-cds-snc-accordion"><summary>([^<]+)<\/summary>\n*(.+)\n*<\/details>/g,
        '<gcds-details details-title="$1">$2</gcds-details>',
      )
      .replace(
        /<div class="wp-block-cds-snc-accordion__content">(.+)<\/div>/g,
        "<gcds-text>$1</gcds-text>",
      )
      .replace(siteUrl, "");
  },
};

export default templateHelpers;
