(function () {
  const htmlElement = document.querySelector("html");
  const colorThemeTogglerButton = document.querySelector(
    "button.theme-toggler"
  );

  if (!colorThemeTogglerButton) {
    console.error("[color-theme-toggler] Missing toggler button.");
    return;
  }

  colorThemeTogglerButton.addEventListener("click", (event) => {
    const currentThemeColor = htmlElement.getAttribute("data-theme");

    if (!currentThemeColor) {
      console.error(
        "[color-theme-toggler] The value for 'currentThemeColor' of the attribute 'data-theme' for the 'html' element is null."
      );
      return;
    }

    switch (currentThemeColor) {
      case "light":
        htmlElement.setAttribute("data-theme", "dark");
        break;

      case "dark":
        htmlElement.setAttribute("data-theme", "light");
        break;

      default:
        console.error(
          "[color-theme-toggler] 'currentThemeColor' has an invalid value."
        );
        return;
    }
  });
})();
