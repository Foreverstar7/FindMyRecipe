import Api from "../../services/Api";

export const autocompleteTags = async(substr) => {
    try {
      const response = await Api.post(`/tags/autocomplete`, `substring=${substr}`);
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const getTags = async() => {
    try {
      const response = await Api.get(`/tags`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
}

export const getTagsOfRecipe = async(recipeid) => {
    try {
      const response = await Api.get(`/tags/recipeid/${recipeid}`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
}

export const getTagByName = async(name) => {
    try {
      const response = await Api.get(`/tags/name/${encodeURIComponent(name)}`, "{}");
      return response;
    } catch (error) {
      console.error(error);
    }
};

