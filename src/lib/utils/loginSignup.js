import Api from "../../services/Api";

export const loginUser = async(username, password) => {
    try {
      const request = JSON.stringify({
        username,
        password
      })
      const response = await Api.post(`/login`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};


export const signupUser = async(username, password) => {
    try {
      const request = JSON.stringify({
        username,
        password
      })
      const response = await Api.post(`/signup`, request, {"Content-type":"application/json"});
      return response;
    } catch (error) {
      console.error(error);
    }
};

export const logout = async() => {
  try {
    const request = JSON.stringify({})
    const response = await Api.post(`/logout`, request, {"Content-type":"application/json"})
    return response
  } catch (error) {
    console.error(error)
  }
}

export const refresh = async(token) => {
  try {
    const request = JSON.stringify(token)
    const response = await Api.post(`/refresh`, request, {"Content-type":"application/json"})
    return response
  } catch (error) {
    console.error(error)
  }
}
