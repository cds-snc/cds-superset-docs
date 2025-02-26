import globals from "globals";
import pluginJs from "@eslint/js";


export default [
  {
    languageOptions: { 
      globals: {
        ...globals.node,
        describe: "readonly",
        jest: "readonly",
        it: "readonly",
        expect: "readonly",
        beforeAll: "readonly",
        beforeEach: "readonly",
        afterAll: "readonly",
        afterEach: "readonly",
        mock: "readonly"
      }
    }
  },
  pluginJs.configs.recommended,
];