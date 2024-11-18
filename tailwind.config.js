const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["internal/views/*.templ", "internal/views/*.go", "internal/views/components/*.templ", "internal/views/components/*.go"],
    theme: {
        container: {
            center: true,
            padding: {
                DEFAULT: "1rem",
                mobile: "2rem",
                tablet: "4rem",
                desktop: "5rem",
            },
        },
        extend: {
            colors: {
                primary: colors.blue,
                secondary: colors.yellow,
                neutral: colors.gray,
            },
        },
    },
    plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};