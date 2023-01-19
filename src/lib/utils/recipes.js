import Api from "../../services/Api";

export const matchRecipes = async(ingredientIds, tagIds) => {
    try {
      const request = JSON.stringify({
        ingredientIds,
        tagIds
      }) 
      const response = await Api.post(`/recipes/match`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};


export const getSimilarRecipes = async(ingredientIds) => {
    try {
      const request = JSON.stringify({
        ingredientIds 
      }) 
      const response = await Api.post(`/recipes/similar`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getUserAuthoredRecipes = async(userId) => {
    try {
      const response = await Api.get(`/recipes/author/${userId}`, {});
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const recommendRecipes = async(userId) => {
    try {
      const request = JSON.stringify({
        userId
      })
      const response = await Api.post(`/recipes/recommend`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getRecipeById = async(id) => {
    try {
      const response = await Api.get(`/recipes/id/${id}`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getRecipes = async() => {
    try {
      const response = await Api.get(`/recipes`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getRecipesForPage = async(page) => {
    try {
      const response = await Api.get(`/recipes?page=${page}`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getWeeklyTopRecipes = async() => {
    try {
      const response = await Api.get(`/recipes/weekly`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};


export const addNewRecipe = async(userId, recipe) => {
    try {
      const request = JSON.stringify({
        userId,
        ...recipe
      })
      const response = await Api.post(`/recipes/add`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};
