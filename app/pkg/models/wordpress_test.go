package models

import (
	"strings"
	"testing"
)

// TestNewPageData tests the NewPageData function which creates page rendering data
func TestNewPageData(t *testing.T) {
	testCases := []struct {
		name              string
		page              WordPressPage
		menu              *MenuData
		googleAnalyticsID string
		siteNames         map[string]string
		baseUrl           string
		expectedData      PageData
	}{
		{
			name: "English page",
			page: WordPressPage{
				ID:       1,
				Slug:     "about",
				SlugEn:   "about",
				SlugFr:   "a-propos",
				Lang:     "en",
				Modified: "2023-05-15T10:30:45",
				Title: struct {
					Rendered string `json:"rendered"`
				}{Rendered: "About Us"},
				Content: struct {
					Rendered string `json:"rendered"`
					Raw      string `json:"raw,omitempty"`
				}{Rendered: "<p>This is content with https://example.com/image.jpg</p>"},
			},
			menu: &MenuData{
				Items: []*MenuItemData{},
			},
			googleAnalyticsID: "UA-12345678-9",
			siteNames: map[string]string{
				"en": "English Site Name",
				"fr": "French Site Name",
			},
			baseUrl: "https://example.com",
			expectedData: PageData{
				GoogleAnalyticsID: "UA-12345678-9",
				Lang:              "en",
				LangSwapPath:      "/fr/",
				LangSwapSlug:      "a-propos",
				Home:              "/",
				Modified:          "2023-05-15",
				Title:             "About Us",
				Content:           "<p>This is content with /image.jpg</p>",
				ShowBreadcrumb:    true,
				SiteName:          "English Site Name",
			},
		},
		{
			name: "French page",
			page: WordPressPage{
				ID:       2,
				Slug:     "a-propos",
				SlugEn:   "about",
				SlugFr:   "a-propos",
				Lang:     "fr",
				Modified: "2023-05-15T10:30:45",
				Title: struct {
					Rendered string `json:"rendered"`
				}{Rendered: "À propos"},
				Content: struct {
					Rendered string `json:"rendered"`
					Raw      string `json:"raw,omitempty"`
				}{Rendered: "<p>C'est du contenu avec https://example.com/image.jpg</p>"},
			},
			menu: &MenuData{
				Items: []*MenuItemData{},
			},
			googleAnalyticsID: "UA-12345678-9",
			siteNames: map[string]string{
				"en": "English Site Name",
				"fr": "French Site Name",
			},
			baseUrl: "https://example.com",
			expectedData: PageData{
				GoogleAnalyticsID: "UA-12345678-9",
				Lang:              "fr",
				LangSwapPath:      "/",
				LangSwapSlug:      "about",
				Home:              "/fr/",
				Modified:          "2023-05-15",
				Title:             "À propos",
				Content:           "<p>C'est du contenu avec /image.jpg</p>",
				ShowBreadcrumb:    true,
				SiteName:          "French Site Name",
			},
		},
		{
			name: "Invalid language defaulting to English",
			page: WordPressPage{
				ID:       3,
				Slug:     "about",
				SlugEn:   "about",
				SlugFr:   "a-propos",
				Lang:     "es", // Invalid language
				Modified: "2023-05-15T10:30:45",
				Title: struct {
					Rendered string `json:"rendered"`
				}{Rendered: "About Us"},
				Content: struct {
					Rendered string `json:"rendered"`
					Raw      string `json:"raw,omitempty"`
				}{Rendered: "<p>Content</p>"},
			},
			menu: &MenuData{
				Items: []*MenuItemData{},
			},
			googleAnalyticsID: "UA-12345678-9",
			siteNames: map[string]string{
				"en": "English Site Name",
				"fr": "French Site Name",
			},
			baseUrl: "https://example.com",
			expectedData: PageData{
				GoogleAnalyticsID: "UA-12345678-9",
				Lang:              "en",
				LangSwapPath:      "/fr/",
				LangSwapSlug:      "a-propos",
				Home:              "/",
				Modified:          "2023-05-15",
				Title:             "About Us",
				Content:           "<p>Content</p>",
				ShowBreadcrumb:    true,
				SiteName:          "English Site Name",
			},
		},
		{
			name: "Home page (no breadcrumb)",
			page: WordPressPage{
				ID:       4,
				Slug:     "home",
				SlugEn:   "home",
				SlugFr:   "accueil",
				Lang:     "en",
				Modified: "2023-05-15T10:30:45",
				Title: struct {
					Rendered string `json:"rendered"`
				}{Rendered: "Home Page"},
				Content: struct {
					Rendered string `json:"rendered"`
					Raw      string `json:"raw,omitempty"`
				}{Rendered: "<p>Welcome home</p>"},
			},
			menu: &MenuData{
				Items: []*MenuItemData{},
			},
			googleAnalyticsID: "UA-12345678-9",
			siteNames: map[string]string{
				"en": "English Site Name",
				"fr": "French Site Name",
			},
			baseUrl: "https://example.com",
			expectedData: PageData{
				GoogleAnalyticsID: "UA-12345678-9",
				Lang:              "en",
				LangSwapPath:      "/fr/",
				LangSwapSlug:      "accueil",
				Home:              "/",
				Modified:          "2023-05-15",
				Title:             "Home Page",
				Content:           "<p>Welcome home</p>",
				ShowBreadcrumb:    false, // Home page, no breadcrumb
				SiteName:          "English Site Name",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the page as we're passing a pointer
			page := tc.page

			// Call the function being tested
			result := NewPageData(&page, tc.menu, tc.googleAnalyticsID, tc.siteNames, tc.baseUrl)

			// Verify results
			if result.GoogleAnalyticsID != tc.expectedData.GoogleAnalyticsID {
				t.Errorf("Expected GoogleAnalyticsID %q, got %q", tc.expectedData.GoogleAnalyticsID, result.GoogleAnalyticsID)
			}

			if result.Lang != tc.expectedData.Lang {
				t.Errorf("Expected Lang %q, got %q", tc.expectedData.Lang, result.Lang)
			}

			if result.LangSwapPath != tc.expectedData.LangSwapPath {
				t.Errorf("Expected LangSwapPath %q, got %q", tc.expectedData.LangSwapPath, result.LangSwapPath)
			}

			if result.LangSwapSlug != tc.expectedData.LangSwapSlug {
				t.Errorf("Expected LangSwapSlug %q, got %q", tc.expectedData.LangSwapSlug, result.LangSwapSlug)
			}

			if result.Home != tc.expectedData.Home {
				t.Errorf("Expected Home %q, got %q", tc.expectedData.Home, result.Home)
			}

			if result.Modified != tc.expectedData.Modified {
				t.Errorf("Expected Modified %q, got %q", tc.expectedData.Modified, result.Modified)
			}

			if string(result.Title) != string(tc.expectedData.Title) {
				t.Errorf("Expected Title %q, got %q", tc.expectedData.Title, result.Title)
			}

			if string(result.Content) != string(tc.expectedData.Content) {
				t.Errorf("Expected Content %q, got %q", tc.expectedData.Content, result.Content)
			}

			if result.ShowBreadcrumb != tc.expectedData.ShowBreadcrumb {
				t.Errorf("Expected ShowBreadcrumb %v, got %v", tc.expectedData.ShowBreadcrumb, result.ShowBreadcrumb)
			}

			if result.SiteName != tc.expectedData.SiteName {
				t.Errorf("Expected SiteName %q, got %q", tc.expectedData.SiteName, result.SiteName)
			}

			// Menu is passed by reference, so it should be the same object
			if result.Menu != tc.menu {
				t.Errorf("Expected Menu to be the same object that was passed in")
			}
		})
	}
}

// TestNewMenuData tests the NewMenuData function which creates hierarchical menu data
func TestNewMenuData(t *testing.T) {
	testCases := []struct {
		name             string
		menuItems        []WordPressMenuItem
		baseUrl          string
		expectedTopItems int
		expectedChildren map[string]int // Map of parent title to number of children
	}{
		{
			name: "Simple menu with no children",
			menuItems: []WordPressMenuItem{
				{
					ID: 1,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "Home"},
					Parent: 0,
					Url:    "https://example.com/",
				},
				{
					ID: 2,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "About"},
					Parent: 0,
					Url:    "https://example.com/about",
				},
			},
			baseUrl:          "https://example.com",
			expectedTopItems: 2,
			expectedChildren: map[string]int{
				"Home":  0,
				"About": 0,
			},
		},
		{
			name: "Menu with parent-child relationships",
			menuItems: []WordPressMenuItem{
				{
					ID: 1,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "Home"},
					Parent: 0,
					Url:    "https://example.com/",
				},
				{
					ID: 2,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "Products"},
					Parent: 0,
					Url:    "https://example.com/products",
				},
				{
					ID: 3,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "Product A"},
					Parent: 2, // Child of Products
					Url:    "https://example.com/products/a",
				},
				{
					ID: 4,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "Product B"},
					Parent: 2, // Child of Products
					Url:    "https://example.com/products/b",
				},
				{
					ID: 5,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "About"},
					Parent: 0,
					Url:    "https://example.com/about",
				},
			},
			baseUrl:          "https://example.com",
			expectedTopItems: 3,
			expectedChildren: map[string]int{
				"Home":     0,
				"Products": 2,
				"About":    0,
			},
		},
		{
			name: "Menu with URL base replacement",
			menuItems: []WordPressMenuItem{
				{
					ID: 1,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "Home"},
					Parent: 0,
					Url:    "https://other-domain.com/",
				},
				{
					ID: 2,
					Title: struct {
						Rendered string `json:"rendered"`
					}{Rendered: "About"},
					Parent: 0,
					Url:    "https://other-domain.com/about",
				},
			},
			baseUrl:          "https://other-domain.com",
			expectedTopItems: 2,
			expectedChildren: map[string]int{
				"Home":  0,
				"About": 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Copy the menu items as we're passing a pointer
			menuItems := tc.menuItems

			// Call the function being tested
			result := NewMenuData(&menuItems, tc.baseUrl)

			// Verify results
			if len(result.Items) != tc.expectedTopItems {
				t.Errorf("Expected %d top-level menu items, got %d", tc.expectedTopItems, len(result.Items))
			}

			// Create a map to easily look up items by title
			itemMap := make(map[string]*MenuItemData)
			for _, item := range result.Items {
				itemMap[item.Title] = item

				// Check if baseUrl is properly replaced in URLs
				if strings.Contains(item.Url, tc.baseUrl) {
					t.Errorf("URL should not contain base URL %q, got %q", tc.baseUrl, item.Url)
				}
			}

			// Check children counts
			for title, expectedCount := range tc.expectedChildren {
				item, exists := itemMap[title]
				if !exists {
					t.Errorf("Expected menu item with title %q not found", title)
					continue
				}

				actualCount := len(item.Children)
				if actualCount != expectedCount {
					t.Errorf("Expected item %q to have %d children, got %d", title, expectedCount, actualCount)
				}

				// For items with children, verify they don't contain the base URL
				for _, child := range item.Children {
					if strings.Contains(child.Url, tc.baseUrl) {
						t.Errorf("Child URL should not contain base URL %q, got %q", tc.baseUrl, child.Url)
					}
				}
			}

			// For nested relationships, check parent-child connections
			if tc.name == "Menu with multiple levels of nesting" {
				productsItem := itemMap["Products"]
				if len(productsItem.Children) != 1 {
					t.Fatalf("Expected Products to have 1 child, got %d", len(productsItem.Children))
				}

				categoryA := productsItem.Children[0]
				if categoryA.Title != "Category A" {
					t.Errorf("Expected child of Products to be 'Category A', got %q", categoryA.Title)
				}

				if len(categoryA.Children) != 1 {
					t.Fatalf("Expected Category A to have 1 child, got %d", len(categoryA.Children))
				}

				productA1 := categoryA.Children[0]
				if productA1.Title != "Product A1" {
					t.Errorf("Expected child of Category A to be 'Product A1', got %q", productA1.Title)
				}
			}
		})
	}
}

// TestConvertToDesignSystem tests the convertToDesignSystem function which transforms
// WordPress HTML content to Design System components
func TestConvertToDesignSystem(t *testing.T) {
	testCases := []struct {
		name         string
		pageContent  string
		baseUrl      string
		expectedHTML string
	}{
		{
			name:         "Base URL removal",
			pageContent:  "This is a link to https://example.com/page and another to https://example.com/other.",
			baseUrl:      "https://example.com",
			expectedHTML: "This is a link to /page and another to /other.",
		},
		{
			name:         "Alert pattern conversion",
			pageContent:  `<details class="alert alert-info" open><summary class="h3"><h3>Information</h3></summary><p>Some important information here.</p></details>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<section class="mt-300 mb-300"><gcds-notice type="info" notice-title-tag="h2" notice-title="Information"><gcds-text><p>Some important information here.</p></gcds-text></gcds-notice></section>`,
		},
		{
			name:         "Button pattern conversion",
			pageContent:  `<div class="wp-block-button"><a class="wp-block-button__link wp-element-button" href="/contact">Contact Us</a></div>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<gcds-button type="link" href="/contact">Contact Us</gcds-button>`,
		},
		{
			name:         "Accordion pattern conversion",
			pageContent:  `<details class="wp-block-cds-snc-accordion"><summary>Frequently Asked Questions</summary><div class="wp-block-cds-snc-accordion__content"><p>FAQ content goes here</p></div></details>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<gcds-details details-title="Frequently Asked Questions"><gcds-text><p>FAQ content goes here</p></gcds-text></gcds-details>`,
		},
		{
			name:         "Accordion content pattern conversion only",
			pageContent:  `<div class="wp-block-cds-snc-accordion__content"><p>Some standalone content</p></div>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<gcds-text><p>Some standalone content</p></gcds-text>`,
		},
		{
			name:         "Multiple patterns in one content",
			pageContent:  `<details class="alert alert-warning" open><summary class="h3"><h3>Warning</h3></summary><p>Be careful!</p></details><div class="wp-block-button"><a class="wp-block-button__link wp-element-button" href="https://example.com/help">Get Help</a></div>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<section class="mt-300 mb-300"><gcds-notice type="warning" notice-title-tag="h2" notice-title="Warning"><gcds-text><p>Be careful!</p></gcds-text></gcds-notice></section><gcds-button type="link" href="/help">Get Help</gcds-button>`,
		},
		{
			name:         "Complex nested patterns",
			pageContent:  `<details class="wp-block-cds-snc-accordion"><summary>FAQs</summary><div class="wp-block-cds-snc-accordion__content"><p>Question 1</p><details class="alert alert-info" open><summary class="h3"><h3>Note</h3></summary><p>Additional info</p></details></div></details>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<gcds-details details-title="FAQs"><gcds-text><p>Question 1</p><section class="mt-300 mb-300"><gcds-notice type="info" notice-title-tag="h2" notice-title="Note"><gcds-text><p>Additional info</p></gcds-text></gcds-notice></section></gcds-text></gcds-details>`,
		},
		{
			name:         "Empty content",
			pageContent:  "",
			baseUrl:      "https://example.com",
			expectedHTML: "",
		},
		{
			name:         "Content with no patterns",
			pageContent:  "<p>This is regular content with no special patterns.</p>",
			baseUrl:      "https://example.com",
			expectedHTML: "<p>This is regular content with no special patterns.</p>",
		},
		{
			name:         "Alert with different types",
			pageContent:  `<details class="alert alert-success" open><summary class="h3"><h3>Success</h3></summary><p>Operation completed successfully.</p></details>`,
			baseUrl:      "https://example.com",
			expectedHTML: `<section class="mt-300 mb-300"><gcds-notice type="success" notice-title-tag="h2" notice-title="Success"><gcds-text><p>Operation completed successfully.</p></gcds-text></gcds-notice></section>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested
			result := convertToDesignSystem(tc.pageContent, tc.baseUrl)

			// Verify result
			if result != tc.expectedHTML {
				t.Errorf("Expected HTML:\n%s\n\nGot:\n%s", tc.expectedHTML, result)
			}
		})
	}
}

// TestConvertToDesignSystemTemplatePreservation tests that the convertToDesignSystem
// function properly preserves template variables and structures
func TestConvertToDesignSystemTemplatePreservation(t *testing.T) {
	// Test case with template variables in content
	pageContent := `<details class="alert alert-info" open><summary class="h3"><h3>{{ .Title }}</h3></summary>{{ .Content }}</details>`
	baseUrl := "https://example.com"
	expected := `<section class="mt-300 mb-300"><gcds-notice type="info" notice-title-tag="h2" notice-title="{{ .Title }}"><gcds-text>{{ .Content }}</gcds-text></gcds-notice></section>`

	result := convertToDesignSystem(pageContent, baseUrl)

	if result != expected {
		t.Errorf("Template variables not preserved correctly.\nExpected: %s\nGot: %s", expected, result)
	}
}

// TestConvertToDesignSystemRealWorldExamples tests the convertToDesignSystem function
// with realistic content examples
func TestConvertToDesignSystemRealWorldExamples(t *testing.T) {
	testCases := []struct {
		name         string
		pageContent  string
		baseUrl      string
		expectedHTML string
	}{
		{
			name: "Page with multiple components",
			pageContent: `<h1>Welcome to our site</h1>
<p>Here's some important information:</p>
<details class="alert alert-info" open><summary class="h3"><h3>Information Notice</h3></summary><p>This service will be down for maintenance on Sunday.</p></details>
<p>Need help? Click the button below:</p>
<div class="wp-block-button"><a class="wp-block-button__link wp-element-button" href="https://example.com/contact">Contact Support</a></div>
<h2>FAQ</h2>
<details class="wp-block-cds-snc-accordion"><summary>How do I reset my password?</summary><div class="wp-block-cds-snc-accordion__content"><p>Visit the account page and click "Reset Password".</p></div></details>`,
			baseUrl: "https://example.com",
			expectedHTML: `<h1>Welcome to our site</h1>
<p>Here's some important information:</p>
<section class="mt-300 mb-300"><gcds-notice type="info" notice-title-tag="h2" notice-title="Information Notice"><gcds-text><p>This service will be down for maintenance on Sunday.</p></gcds-text></gcds-notice></section>
<p>Need help? Click the button below:</p>
<gcds-button type="link" href="/contact">Contact Support</gcds-button>
<h2>FAQ</h2>
<gcds-details details-title="How do I reset my password?"><gcds-text><p>Visit the account page and click "Reset Password".</p></gcds-text></gcds-details>`,
		},
		{
			name: "Edge case with partial matches",
			pageContent: `<p>This is a normal paragraph</p>
<div class="not-wp-block-button"><a href="/link">Not a button</a></div>
<details><summary>Not an accordion</summary>Some content</details>
<div class="something-else">Not accordion content</div>`,
			baseUrl: "https://example.com",
			expectedHTML: `<p>This is a normal paragraph</p>
<div class="not-wp-block-button"><a href="/link">Not a button</a></div>
<details><summary>Not an accordion</summary>Some content</details>
<div class="something-else">Not accordion content</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := convertToDesignSystem(tc.pageContent, tc.baseUrl)

			if result != tc.expectedHTML {
				t.Errorf("Content transformation failed.\nExpected:\n%s\n\nGot:\n%s", tc.expectedHTML, result)
			}
		})
	}
}
