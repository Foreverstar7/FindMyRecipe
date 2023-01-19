export const NUTRITION_NAMES = [
  'Calories',
  'Total fat (%DV)',
  'Sugar (%DV)',
  'Sodium (%DV)',
  'Protein (%DV)',
  'Saturated Fat (%DV)',
  'Carbohydrates (%DV)',
]

export function capitalizeFirstChar(text) {
  return text[0].toUpperCase() + text.slice(1)
}

export const ENTER_KEY = 13

export const RECIPE_NUMS = 231637 // default unless adding recipe is a thing

export const USER_FEATURES = [
  'Add your favorite recipes!',
  'Get recommended recipes!',
  'Write your own recipes!',
]

export const getAuthToken = () => {
  const jwtToken =
    document.cookie.match('(^|;)\\s*' + 'token' + '\\s*=\\s*([^;]+)')?.pop() ||
    ''
  if (jwtToken == '') {
    return {}
  }
  let base64Url = jwtToken.split('.')[1]
  let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
  let jsonPayload = decodeURIComponent(
    window
      .atob(base64)
      .split('')
      .map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
      })
      .join(''),
  )
  return JSON.parse(jsonPayload)
}

export const getUserInfo = () => {
  const authToken = getAuthToken()
  if (
    authToken.hasOwnProperty('username') &&
    authToken.hasOwnProperty('userid') &&
    authToken.hasOwnProperty('exp')
  ) {
    return authToken
  }
  return null
}
