```js
// Grab the display box
const display = document.getElementById("display");

// Track the full expression typed by user
let expression = "";

// Handle all button clicks
document.querySelectorAll("button").forEach(button => {
    button.addEventListener("click", () => {
        const value = button.textContent;

        if (value === "C") {
            expression = "";
            display.value = "";
        }
        else if (value === "=") {
            // Send expression to Go backend
            fetch("/api/calc", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ expr: expression })
            })
            .then(res => res.json())
            .then(data => {
                display.value = data.result;
                expression = data.result.toString();
            })
            .catch(() => {
                display.value = "mperi";
                expression = "";
            });
        }
        else {
            expression += value;
            display.value = expression;
        }
    });
});
```