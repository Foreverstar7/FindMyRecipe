import Api from "../../services/Api";

export const autocompleteIngredients = async(substr) => {
    try {
      const response = await Api.post(`/ingredients/autocomplete`, `substring=${substr}`);
      return response;
    } catch (error) {
      console.error(error);
    }
};


export const getIngredientById = async(id) => {
    try {
      const response = await Api.get(`/ingredients/id/${id}`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getIngredientByName = async(name) => {
    try {
      const response = await Api.get(`/ingredients/name/${encodeURIComponent(name)}`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getIngredientsOfRecipe = async(id) => {
    try {
      const response = await Api.get(`/ingredients/byRecipeId/${id}`, "{}");
      return response;
    } catch (error) {
      console.error(error)
    }
}
