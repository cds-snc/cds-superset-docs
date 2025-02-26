import { escapeRegExp } from "../../utils/string-helpers";

describe("utils/string-helpers", () => {
  describe("escapeRegExp", () => {
    it("should return same string when no special characters present", () => {
      expect(escapeRegExp("hello")).toBe("hello");
      expect(escapeRegExp("test123")).toBe("test123");
    });

    it("should escape single special characters", () => {
      expect(escapeRegExp(".")).toBe("\\.");
      expect(escapeRegExp("*")).toBe("\\*");
      expect(escapeRegExp("?")).toBe("\\?");
    });

    it("should escape multiple special characters", () => {
      expect(escapeRegExp("hello.world")).toBe("hello\\.world");
      expect(escapeRegExp("test*string")).toBe("test\\*string");
    });

    it("should escape common regex patterns", () => {
      expect(escapeRegExp("[0-9]+")).toBe("\\[0-9\\]\\+");
      expect(escapeRegExp("(test)")).toBe("\\(test\\)");
    });

    it("should handle empty string", () => {
      expect(escapeRegExp("")).toBe("");
    });

    it("should handle undefined", () => {
      expect(escapeRegExp(undefined)).toBe(undefined);
    });

    it("should escape all special regex characters", () => {
      const specialChars = ".*+?^${}()|[]\\";
      const expected = "\\.\\*\\+\\?\\^\\$\\{\\}\\(\\)\\|\\[\\]\\\\";
      expect(escapeRegExp(specialChars)).toBe(expected);
    });
  });
});
