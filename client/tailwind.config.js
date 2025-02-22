/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      "./index.html",
      "./src/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sfpro: ["SF-Pro", "system-ui"]
      },
      height: {
        "1/10": "10%",
        "9/10": "90%"
      },
      colors: {
        mainbg: "#161616",
        lightgray: "#B3B3B3",
        midgray: "#6E6E6E",
        darkgray: "#575757",
        searchbargray: "#222222",
        searchbarbg: "#1D1D1D",
        grad1: "#29A468",
        grad2: "#92C066",
        grad3: "#9AC069",
        grad4: "#6CC06F",
        grad5: "#46B7B9",
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}

