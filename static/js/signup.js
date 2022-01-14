// Signup Page Functionality
(function () {
    let password = document.getElementById("password");
    let confirmPassword = document.getElementById("confirmPassword");

    let matchingMessage = document.getElementById("matchingMessage");
    let passwordInstructions = document.getElementById("passwordInstructions");

    let letter = document.getElementById("letter");
    let capital = document.getElementById("capital");
    let number = document.getElementById("number");
    let passLength = document.getElementById("length");
    let symbol = document.getElementById("symbol");

    confirmPassword.addEventListener("input", function () {
        if (password.value === confirmPassword.value) {
            matchingMessage.style.borderColor = "rgb(0, 255, 0)";
            matchingMessage.classList.add("hidden");
        } else {
            matchingMessage.style.borderColor = "red";
            matchingMessage.classList.remove("hidden");
            matchingMessage.innerHTML =
                "<p class='form-text'>Passwords don't match</p>";
        }
    });

    // When the user clicks on the password field, show the message box
    password.addEventListener("focus", function () {
        document.getElementById("passwordInstructions").style.display = "block";
    });

    // When the user clicks outside of the password field, hide the message box
    password.addEventListener("blur", function () {
        document.getElementById("passwordInstructions").style.display = "none";
    });

    // When the user starts to type something inside the password
    password.addEventListener("keyup", function () {
        // Validate lowercase letters
        var lowerCaseLetters = /[a-z]/g;
        if (password.value.match(lowerCaseLetters)) {
            letter.classList.remove("invalid");
            letter.classList.add("valid");
        } else {
            letter.classList.remove("valid");
            letter.classList.add("invalid");
        }

        // Validate capital letters
        var upperCaseLetters = /[A-Z]/g;
        if (password.value.match(upperCaseLetters)) {
            capital.classList.remove("invalid");
            capital.classList.add("valid");
        } else {
            capital.classList.remove("valid");
            capital.classList.add("invalid");
        }

        // Validate numbers
        var numbers = /[0-9]/g;
        if (password.value.match(numbers)) {
            number.classList.remove("invalid");
            number.classList.add("valid");
        } else {
            number.classList.remove("valid");
            number.classList.add("invalid");
        }

        // Validate length
        if (password.value.length >= 8) {
            passLength.classList.remove("invalid");
            passLength.classList.add("valid");
        } else {
            passLength.classList.remove("valid");
            passLength.classList.add("invalid");
        }

        // Validate symbol
        var symbols = /[#?!@$ %^&*-]/g;
        if (password.value.match(symbols)) {
            symbol.classList.remove("invalid");
            symbol.classList.add("valid");
        } else {
            symbol.classList.remove("valid");
            symbol.classList.add("invalid");
        }
    });
})();