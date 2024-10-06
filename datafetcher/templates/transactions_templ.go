// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/esteanes/expense-manager/datafetcher/functions"
	"github.com/esteanes/expense-manager/datafetcher/upclient"
	"strconv"
	"time"
)

func getMessage(transaction upclient.TransactionResource) string {
	maybeString := transaction.Attributes.Message.Get()
	if maybeString == nil {
		return ""
	}
	return *maybeString
}

func getCardPurchase(transaction upclient.TransactionResource) string {
	maybeCardPurchase := transaction.Attributes.CardPurchaseMethod.Get()
	if maybeCardPurchase == nil {
		return ""
	}
	return string(maybeCardPurchase.Method)
}

func getCardPurchaseNumber(transaction upclient.TransactionResource) string {
	maybeCardPurchase := transaction.Attributes.CardPurchaseMethod.Get()
	if maybeCardPurchase == nil {
		return ""
	}
	maybePurchaseNumber := maybeCardPurchase.CardNumberSuffix.Get()
	if maybePurchaseNumber == nil {
		return ""
	}
	return *maybePurchaseNumber
}

func getCategory(transaction upclient.TransactionResource) string {
	maybeCategory := transaction.GetRelationships().Category.Data.Get()
	if maybeCategory == nil {
		return ""
	}
	return maybeCategory.Id

}

func getDateOrDefault(dateTime *time.Time, defaultTime time.Time) string {
	if dateTime == nil {
		return defaultTime.Format("2006-01-02")
	}
	return dateTime.Format("2006-01-02")
}

func TransactionsTable(transactions chan upclient.TransactionResource) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto p-8\"><h1 class=\"text-4xl font-bold mb-6\">Transactions</h1><button id=\"downloadButton\" class=\"px-6 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 flex items-center justify-center\"><span id=\"buttonText\">Download Content</span> <svg id=\"loadingSpinner\" class=\"hidden animate-spin h-5 w-5 ml-2 text-white\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\"><circle class=\"opacity-25\" cx=\"12\" cy=\"12\" r=\"10\" stroke=\"currentColor\" stroke-width=\"4\"></circle> <path class=\"opacity-75\" fill=\"currentColor\" d=\"M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z\"></path></svg></button><div class=\"overflow-y-auto max-h-96 border border-gray-200 rounded-lg\"><table class=\"min-w-full table-auto bg-gray-50 border border-gray-200 p-10\"><thead class=\"bg-gray-200\"><tr><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Transaction Amount</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Description</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Status</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Time</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Message</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Category</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Purchase Method</th><th class=\"bg-gray-200 px-4 py-2 text-left border-b border-gray-300 sticky top-0\">Card Number</th></tr></thead> <tbody class=\"bg-grey-light\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for transaction := range transactions {
			templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"hover:bg-yellow-400\"><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(transaction.Attributes.Amount.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 82, Col: 92}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(transaction.Attributes.Description)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 83, Col: 91}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(string(transaction.Attributes.Status))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 84, Col: 94}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(transaction.Attributes.CreatedAt.String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 85, Col: 98}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(getMessage(transaction))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 86, Col: 80}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(getCategory(transaction))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 87, Col: 81}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(getCardPurchase(transaction))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 88, Col: 85}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-2 border-b border-gray-300\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(getCardPurchaseNumber(transaction))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 89, Col: 91}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = templ.Flush().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div></div><script>\n        document.getElementById('downloadButton').addEventListener('click', function() {\n            const startDateInput = document.getElementById('start-date');\n            const endDateInput = document.getElementById('end-date');\n            const startDate = new Date(startDateInput.value).toISOString();\n            const endDate = new Date(endDateInput.value).toISOString();\n            const numTransactions = document.getElementById('slider').value;\n\n            // Build the URL with query parameters\n            const url = new URL('/transactions-csv', window.location.origin);\n            if (startDate) {\n                url.searchParams.set('startDate', startDate);\n            }\n            if (endDate) {\n                url.searchParams.set('endDate', endDate);\n            }\n            url.searchParams.set('numTransactions', numTransactions);\n            url.searchParams.set('accountId', new URLSearchParams(window.location.search).get(\"accountId\"))\n\n            // Fetch the file with the generated URL\n            fetch(url, {\n                method: 'GET',\n                headers: {\n                    // Add any headers if necessary\n                }\n            })\n            .then(response => response.blob())  // Convert to blob\n            .then(blob => {\n                // Create a URL for the blob object\n                const downloadUrl = window.URL.createObjectURL(blob);\n\n                // Create a temporary <a> element for the download\n                const a = document.createElement('a');\n                a.href = downloadUrl;\n                a.download = 'upbank-downloads.csv';  // Specify the filename\n                document.body.appendChild(a);\n\n                // Trigger the download\n                a.click();\n\n                // Clean up\n                a.remove();\n                window.URL.revokeObjectURL(downloadUrl);\n            })\n            .catch(error => console.error('Download failed', error));\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func numTransactionSlider(numTransactions string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"p-6 bg-white rounded-lg shadow-lg\"><label for=\"slider\" class=\"block text-lg font-semibold mb-4\">Number of Transactions</label><!-- Slider --><input type=\"range\" id=\"slider\" min=\"1\" max=\"10000\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(numTransactions)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 150, Col: 77}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full h-2 bg-gray-300 rounded-lg appearance-none cursor-pointer\"><!-- Displayed value as an input field --><div class=\"mt-4 flex items-center space-x-2\"><label for=\"sliderValue\" class=\"text-sm font-medium text-gray-700\">Value:</label> <input type=\"number\" id=\"sliderValue\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(numTransactions)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 154, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"1\" max=\"100\" class=\"w-32 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500\"></div><!-- Button to grab the slider value and refresh the page --><button id=\"submitButton\" class=\"mt-6 px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400\">Set Transactions</button></div><script>\n        const slider = document.getElementById('slider');\n        const sliderValueInput = document.getElementById('sliderValue');\n        const submitButton = document.getElementById('submitButton');\n\n        // Sync input field when slider changes\n        slider.addEventListener('input', function() {\n            sliderValueInput.value = slider.value;\n        });\n\n        // Sync slider when input field changes\n        sliderValueInput.addEventListener('input', function() {\n            if (sliderValueInput.value >= slider.min && sliderValueInput.value <= slider.max) {\n                slider.value = sliderValueInput.value;\n            }\n        });\n\n        // On button click, update URL with the query parameter and refresh the page\n        submitButton.addEventListener('click', function() {\n            const sliderValue = slider.value;\n            const url = new URL(window.location.href);\n            url.searchParams.set('numTransactions', sliderValue);\n            window.location.href = url.toString(); // Refresh the page with the updated query parameter\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func dateSelector(startDate string, endDate string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"p-6 bg-white rounded-lg shadow-lg max-w-2xl\"><h2 class=\"text-xl font-bold mb-4 text-gray-700\">Select Date Range</h2><div class=\"flex space-x-4\"><!-- Start Date Input --><div class=\"flex flex-col\"><label for=\"start-date\" class=\"mb-2 text-sm font-medium text-gray-600\">Start Date</label> <input type=\"date\" id=\"start-date\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(startDate)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 195, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500\"></div><!-- End Date Input --><div class=\"flex flex-col\"><label for=\"end-date\" class=\"mb-2 text-sm font-medium text-gray-600\">End Date</label> <input type=\"date\" id=\"end-date\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(endDate)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `datafetcher/templates/transactions.templ`, Line: 200, Col: 52}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500\"></div></div><button id=\"dateSubmitButton\" class=\"mt-4 px-6 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400\">Submit</button><script>\n        // On button click, update URL with the query parameters for startDate and endDate and refresh the page\n        dateSubmitButton.addEventListener('click', function() {\n            const startDateInput = document.getElementById('start-date');\n            const endDateInput = document.getElementById('end-date');\n            const dateSubmitButton = document.getElementById('dateSubmitButton');\n            const startDate = new Date(startDateInput.value).toISOString();\n            const endDate = new Date(endDateInput.value).toISOString();\n\n            // Create a new URL object\n            const url = new URL(window.location.href);\n\n            // Add or update the query parameters for startDate and endDate\n            if (startDate) {\n                url.searchParams.set('startDate', startDate);\n            }\n            if (endDate) {\n                url.searchParams.set('endDate', endDate);\n            }\n            if (startDate == null && endDate == null) {\n                return;\n            }\n\n            // Refresh the page with the updated URL\n            window.location.href = url.toString();\n        });\n    </script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Transactions(title string, transactions chan upclient.TransactionResource, accounts chan upclient.AccountResource, queryParams *functions.QueryParams) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var18 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Var19 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				templ_7745c5c3_Err = AccountButtons(accounts, false).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = numTransactionSlider(strconv.Itoa(int(*queryParams.NumTransactions))).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = dateSelector(getDateOrDefault(queryParams.StartDate, time.Now().AddDate(0, 0, -1)), getDateOrDefault(queryParams.EndDate, time.Now())).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = GridOrganiser().Render(templ.WithChildren(ctx, templ_7745c5c3_Var19), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = TransactionsTable(transactions).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base(title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var18), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
