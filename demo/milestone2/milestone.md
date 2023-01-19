# Milestone 2: Features

- Make sure to unzip the production data in data/data.zip
- Code to initialize API
```
go mod tidy
demo/milestone2/milestone2.sh
demo/milestone2/populateSample.sh // For sample data
demo/milestone2/populateProd.sh // For production data
go run .
```

## Login and Sign Up
- Sign Up Endpoint
```
curl -d '{"username": "testfeature", "password": "testpass"}' http://localhost:8080/api/signup/
```
- Expected Output
```
{ "status": "success" } // If user successfully registered
{ "status": "failed", "message": "username already exists" } 
```

- Login Endpoint
```
curl -d '{"username": "testfeature", "password": "testpass"}' http://localhost:8080/api/login/
```
- Expected Output
```
{ "status": "success" } // If username and password matches
{ "status": "failed" } // If username and password don't match
```

## Autocomplete Ingredients
- Autocomplete Endpoint
```
curl -d 'substring=bac' http://localhost:8080/api/ingredients/autocomplete/
```
- Expected Output (Sample Data):
```
{"data":[{"id":11,"name":"bacon"}],"message":"List of Similar Ingredients","status":"success"}
```
- Expected Output (Production Data):
```
{"data":[{"id":294,"name":"tomato bacon salad dressing"},{"id":327,"name":"epicure cheese chives and bacon dip mix"},{"id":453,"name":"bacardi dark rum"},{"id":489,"name":"side bacon"},{"id":553,"name":"imitation bacon bits"},{"id":739,"name":"backfin crab meat"},{"id":760,"name":"maple bacon"},{"id":953,"name":"pork fatback"},{"id":1128,"name":"rack of baby-back pork ribs"},{"id":1420,"name":"back bacon"},{"id":1636,"name":"peameal bacon"},{"id":1828,"name":"turkey bacon"},{"id":1967,"name":"vegan bacon"},{"id":2287,"name":"vegetarian bacon bits"},{"id":2343,"name":"bacon ranch dressing"},{"id":2382,"name":"louis rich turkey bacon"},{"id":2574,"name":"smoked bacon"},{"id":2651,"name":"streaky bacon"},{"id":2743,"name":"chicken backs and necks"},{"id":2746,"name":"albacore tuna"},{"id":2835,"name":"bacardi 151 rum"},{"id":2843,"name":"bacon bits"},{"id":2845,"name":"pork baby back ribs"},{"id":2849,"name":"bacon"},{"id":2952,"name":"apple-smoked bacon"},{"id":2986,"name":"zwieback toast crumbs"},{"id":3186,"name":"bacon drippings"},{"id":3312,"name":"baby back ribs"},{"id":3882,"name":"bacardi limon"},{"id":4015,"name":"condensed bean with bacon soup"},{"id":4092,"name":"morningstar farms veggie bacon strips"},{"id":4159,"name":"hot bacon salad dressing"},{"id":4426,"name":"oscar mayer bacon"},{"id":4679,"name":"fried bacon"},{"id":4873,"name":"j\u0026d peppered bacon salt"},{"id":5052,"name":"bacon fat"},{"id":5166,"name":"bacon piece"},{"id":5243,"name":"oscar mayer real bacon bits"},{"id":5574,"name":"crisp bacon"},{"id":5631,"name":"chicken bacon"},{"id":5666,"name":"baby back rib rack"},{"id":5739,"name":"pork back ribs"},{"id":5895,"name":"hickory smoked bacon"},{"id":6135,"name":"low-sodium bacon"},{"id":6337,"name":"bacardi peach red rum"},{"id":6378,"name":"vegetarian canadian bacon"},{"id":6584,"name":"rindless bacon"},{"id":6642,"name":"reduced-sodium bacon"},{"id":6876,"name":"pepper bacon"},{"id":7298,"name":"center-cut bacon"},{"id":7352,"name":"bacon dip"},{"id":7385,"name":"lean bacon"},{"id":7505,"name":"bacon grease"},{"id":7533,"name":"bacon cheddar cheese"},{"id":7559,"name":"chicken back"},{"id":7588,"name":"fatty bacon"},{"id":7703,"name":"tempeh bacon"},{"id":8071,"name":"albacore tuna in water"},{"id":8129,"name":"low fat bacon bits"},{"id":8321,"name":"fat free bacon"},{"id":8357,"name":"smoked back bacon"},{"id":8434,"name":"kraft cheese spread with bacon"},{"id":8636,"name":"cooked bacon"},{"id":8783,"name":"beef back ribs"},{"id":8811,"name":"baby back rib racks"},{"id":9068,"name":"fat-free turkey bacon"},{"id":9513,"name":"beef bacon"},{"id":9538,"name":"imitation bacon"},{"id":9583,"name":"rindless smoked streaky bacon"},{"id":9642,"name":"instant mashed potatoes with bacon and cheddar"},{"id":9704,"name":"bean with bacon soup"},{"id":9931,"name":"smoked streaky bacon"},{"id":10123,"name":"bacos bacon bits"},{"id":10289,"name":"sour cream and bacon salad dressing"},{"id":10569,"name":"reduced-fat bacon"},{"id":10750,"name":"cooked turkey bacon"},{"id":10922,"name":"unsmoked bacon"},{"id":11301,"name":"racks of baby-back pork ribs"},{"id":11615,"name":"back ribs"},{"id":11646,"name":"soy bacon bits"},{"id":12055,"name":"bacon pieces"},{"id":12545,"name":"low-fat bacon"},{"id":12548,"name":"zwieback toast"},{"id":12689,"name":"bacardi light rum"},{"id":12715,"name":"hormel real bacon bits"},{"id":12766,"name":"canadian bacon"},{"id":12884,"name":"shortcut bacon"},{"id":12907,"name":"lamb backstraps"},{"id":13105,"name":"thick slab bacon"},{"id":13648,"name":"shoulder bacon"},{"id":13725,"name":"maple cured bacon baked beans"},{"id":13777,"name":"lean beef bacon"},{"id":14164,"name":"low-fat bacon ranch dressing"},{"id":14195,"name":"lamb backstrap"},{"id":14442,"name":"bacardi o rum"},{"id":14486,"name":"bacardi fuzzy navel mix"},{"id":14530,"name":"vegetarian bacon"}],"message":"List of Similar Ingredients","status":"success"}
```

## Matching Recipes
- Recipe Match Endpoint
```
curl -H 'Content-type: application/json' \
  -d '{"ingredientIds": [4621,4684,6557,8899,10583,12245,14832 ]}' \
  http://localhost:8080/api/recipes/match/
```
- Expected Output (Sample Data):
```
{"data":[{"Id":1,"Name":"arriba   baked winter squash mexican style","Description":"autumn is my favorite time of year to cook! this recipe \r\ncan be prepared either spicy or sweet, your choice!\r\ntwo of my posted mexican-inspired seasoning mix recipes are offered as suggestions.","PrepMinutes":55,"Author":1,"Steps":["'make a choice and proceed with recipe'","'depending on size of squash ; cut into half or fourths'","'remove seeds'","'for spicy squash ; drizzle olive oil or melted butter over each cut squash piece'","'season with mexican seasoning mix ii'","'for sweet squash ; drizzle melted honey ; butter ; grated piloncillo over each cut squash piece'","'season with sweet mexican spice mix'","'bake at 350 degrees ; again depending on size ; for 40 minutes up to an hour ; until a fork can easily pierce the skin'","'be careful not to burn the squash especially if you opt to use sugar or butter'","'if you feel more comfortable ; cover the squash with aluminum foil the first half hour ; give or take ; of baking'","'if desired ; season with salt'"],"Nutrition":[51.5,0,13,0,2,0,4]}], "message":"Matching Recipes", "status":"success"}
```
- Expected Output (Production Data):
```
{"data":[{"Id":1,"Name":"arriba   baked winter squash mexican style","Description":"autumn is my favorite time of year to cook! this recipe \r\ncan be prepared either spicy or sweet, your choice!\r\ntwo of my posted mexican-inspired seasoning mix recipes are offered as suggestions.","PrepMinutes":55,"Author":1,"Steps":["'make a choice and proceed with recipe'","'depending on size of squash ; cut into half or fourths'","'remove seeds'","'for spicy squash ; drizzle olive oil or melted butter over each cut squash piece'","'season with mexican seasoning mix ii'","'for sweet squash ; drizzle melted honey ; butter ; grated piloncillo over each cut squash piece'","'season with sweet mexican spice mix'","'bake at 350 degrees ; again depending on size ; for 40 minutes up to an hour ; until a fork can easily pierce the skin'","'be careful not to burn the squash especially if you opt to use sugar or butter'","'if you feel more comfortable ; cover the squash with aluminum foil the first half hour ; give or take ; of baking'","'if desired ; season with salt'"],"Nutrition":[51.5,0,13,0,2,0,4]},{"Id":107268,"Name":"honey butter 101","Description":"honey butter is very good on homemade white bread. i ate and made this many times as a kid. you may also use peanut butter instead of butter. delicious!","PrepMinutes":20,"Author":1,"Steps":["'1'","'place butter in a small bowl'","'gradually add honey ; beating constantly ; until desired thickness is attained'","'my'","'i used this on bagels ; but can definitely see a lot of uses for this stuff and add-ins like orange zest'","'but this is great on homemade dinner rolls ! dinner guests always rave about this butter'","'for easier mixing ; be sure your butter and honey are room temperature also ; i like to use a bit of salted butter because it really does enhance the honey / butter flavor - i suggest half salted and half unsalted'","'i tried just adding salt once ; but it doesnt really disolve well ; so go with the presalted butter'","'a bee-keeper kept bees on our land and paid us in honey so we used honey in everything -- but this was our favorite'"],"Nutrition":[166.1,17,69,4,0,36,5]},{"Id":107272,"Name":"honey butter shanachie","Description":"this was a demonstration recipe at charlotte anne albertson's cooking school (wynnewood, pa), for \"awesome soups with brian duffy\" (jan 24, 2006). brian duffy (http://www.chefduff.com) hosted the food network shows \"date plate\" and \"hot trends 2005\" and is executive chef at shanachie, an irish restaurant in ambler, pa. \r\n\r\nshanachie specializes in soups and serves wheaten bread (rz. 28076) with honey butter with its soups.\r\n\r\nthe original recipe (restaurant quantity) is 30 lbs butter (preferably irish unsalted butter) and 10 lbs honey (preferably clover honey). if you're not feeding an army, this is the reduced recipe, approximately the same proportion ...\r\n\r\nthe \"cooking time\" listed is to warm the butter and then chill it back.\r\n\r\nassuming a 1 tsp portion, this will make 24 pats of butter.","PrepMinutes":25,"Author":1,"Steps":["'warm the butter -- dont need to melt it'","'add the honey to the butter'","'form into a block and chill to reform the shape'","'alternately ; you can form a log ; chill into molds ; cut pats ; etc'"],"Nutrition":[150,23,15,5,0,48,1]},{"Id":107273,"Name":"honey butter spread","Description":"this is a quick and easy recipe for honey butter spread.  i made an elegant luncheon for one of my friends and needed a little something extra to add a special touch.  she took one bite and instantly asked, \"what's in this?  it's so good!\"  i used salted butter because i liked the subtle combination of sweet and salty, but one could use unsalted if desired.  it is best served at room temperature spread on the bread of your choice.","PrepMinutes":5,"Author":1,"Steps":["'place butter and honey in a small mixing bowl'","'with a hand-held mixer ; beat at low speed until well-combined and creamy in texture'","'with a rubber spatula ; scrape butter out of mixing bowl into a serving bowl'","'serve at room temperature'"],"Nutrition":[94.2,14,13,2,0,29,1]},{"Id":107509,"Name":"honey hair conditioner","Description":"add a beautiful  shine to your hair","PrepMinutes":32,"Author":1,"Steps":["'using a small amount at a time work mixture through you hair'","'cover with a shower cap or cling wrap ; let sit 30 minutes'","'shampoo well and rinse thoroughly'","'dry and style as usual'"],"Nutrition":[1323.5,110,742,0,1,49,62]},{"Id":204793,"Name":"susan s honey bee butter","Description":"since we have our own beehives, i'm always on the look-out for great recipes to use up our beautiful wild blend of honey which we harvest/extract every october.  this recipe from susan branch's \"autumn\" book makes cornbread, scones and panny cakes sing with delight!   and you will, too!","PrepMinutes":2,"Author":1,"Steps":["'mix by hand or in a small food chopper'","'youre done !'"],"Nutrition":[117.7,17,17,3,0,36,1]},{"Id":225747,"Name":"whipped honey butter with varieties","Description":"a basic honey butter that's delicious on cornbread or muffins! also, a number of variations. try it on scones!","PrepMinutes":10,"Author":1,"Steps":["'whip butter with mixer until fluffy'","'gradually add in honey and mix until well blended'","'chill until ready to use'","'vanilla cinnamon honey butter: add 1 / 4 teaspoon vanilla and 1 / 4 teaspoon cinnamon'","'orange honey butter: add 1-2 teaspoons grated orange zest'","'lemon honey butter: add 1-2 teaspoons grated lemon zest'"],"Nutrition":[2142.9,283,557,67,4,583,46]}],"message":"Match Recipes by Ingredients","status":"success"}
```

## Filter by Tags
- Tag Filter Endpoint
```
curl -d 'tagId=124' http://localhost:8080/api/recipes/tags
```
- Expected Output
```
{"data":[...],"message":"List of Recipes by Tags","status":"success"}
```

## Get Ingredients by ID
- Ingredient ID Endpoint
```
curl http://localhost:8080/api/ingredients/id/10
```
- Expected Output
```
{"data":[...],"message":"Ingredient by ID","status":"success"}
```

## Get Recipes by ID
- Recipe ID Endpoint
```
curl http://localhost:8080/api/recipes/id/10
```
- Expected Output
```
{"data":[...],"message":"Recipe by ID","status":"success"}
```

## Get Weekly Top Recipes
- Endpoint
```
curl http://localhost:8080/api/recipes/weekly
```
- Expected Output
```
{"data":[...],"message":"Weekly Favorite Recipes","status":"success"}
{"data":[],"message":"No Weekly Favorites","status":"failed"}
```

## Similar Recipes
- Endpoint
```
curl -H 'Content-type: application/json' \
  -d '{"ingredientIds": [4621,4684,6557,8899,10583,12245,14832]}' \
  http://localhost:8080/api/recipes/similar/
```
- Expected Output
```
{"data":[...],"message":"Similar Recipes by Ingredients","status":"success"}
{"data":[],"message":"No Similar Recipes by Ingredients","status":"failed"}
```

## Recommended Recipes for User
- Endpoint
```
curl -d 'userId=1' http://localhost:8080/api/recipes/recommend/
```
- Expected Output
```
{"data":[...],"message":"Recipe Recommendation","status":"success"}
{"data":[],"message":"No Recommendation for User","status":"failed"}
```

