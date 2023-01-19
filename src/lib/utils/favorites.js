import Api from "../../services/Api";

export const toggleUserFavorite = async(userId, recipeId) => {
    try {
      const request = JSON.stringify({
        userId,
        recipeId
      })
      const response = await Api.post(`/favorite/toggle`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getUserFavorites = async(userId) => {
    try {
      const request = JSON.stringify({
        userId
      })
      const response = await Api.post(`/favorite/user`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};

