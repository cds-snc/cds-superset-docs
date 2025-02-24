import templateHelpers from "../../utils/template-helpers";

describe("utils/template-helpers", () => {
  describe("dateFormat", () => {
    it("should format date correctly", () => {
      const date = "2023-09-20T12:00:00Z";
      const result = templateHelpers.dateFormat(date);
      expect(result).toBe("2023-09-20");
    });

    it("should handle invalid date", () => {
      const date = "asdf";
      const result = templateHelpers.dateFormat(date);
      // Your formatting might return "NaN-NaN-NaN" or local fallback
      expect(result).toBe("Invalid Date");
    });
  });

  describe("eq", () => {
    it("should return true if strings match", () => {
      expect(templateHelpers.eq("test", "test")).toBe(true);
    });

    it("should return false if strings do not match", () => {
      expect(templateHelpers.eq("test", "nope")).toBe(false);
    });
  });

  describe("updateMarkup", () => {
    it("should return empty string if content is undefined", () => {
      const result = templateHelpers.updateMarkup(undefined);
      expect(result).toBe("");
    });

    it("should return empty string if content is empty", () => {
      const result = templateHelpers.updateMarkup("");
      expect(result).toBe("");
    });

    it("should transform alert details into gcds-notice sections", () => {
      const content = `<details class="alert alert-warning" open><summary class="h3"><h3>Alert Title</h3></summary>Alert body</details>`;
      const result = templateHelpers.updateMarkup(content);
      expect(result).toContain(
        '<section class="mt-300 mb-300"><gcds-notice type="warning"',
      );
      expect(result).toContain('notice-title="Alert Title"');
      expect(result).toContain("<gcds-text>Alert body</gcds-text>");
    });

    it("should transform WP block button into gcds-button", () => {
      const content = `<div class="wp-block-button"><a class="wp-block-button__link another-class" href="https://example.com">Click me</a></div>`;
      const result = templateHelpers.updateMarkup(content);
      expect(result).toBe(
        `<gcds-button type="link" href="https://example.com">Click me</gcds-button>`,
      );
    });

    it("should transform accordion details into gcds-details", () => {
      const content = `<details class="wp-block-cds-snc-accordion"><summary>Test</summary>\nSome accordion content\n</details>`;
      const result = templateHelpers.updateMarkup(content);
      expect(result).toBe(
        `<gcds-details details-title="Test">Some accordion content</gcds-details>`,
      );
    });

    it("should wrap accordion content in gcds-text", () => {
      const content = `<div class="wp-block-cds-snc-accordion__content">More content here</div>`;
      const result = templateHelpers.updateMarkup(content);
      expect(result).toBe("<gcds-text>More content here</gcds-text>");
    });
  });
});
