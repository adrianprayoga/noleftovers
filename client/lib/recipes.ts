import axios from "axios";

export const getAllRecipes = async (): Promise<any[]> => {
  try {
    console.log("calling", `${process.env.HOST}/recipe`);
    const response = await axios.get(`${process.env.HOST}/recipe`);

    return response.data;
  } catch (err) {
    console.error("error in getting recipes");

    return [];
  }
};

export const getAllRecipeIds = async () => {
  const recipeList = await getAllRecipes()

  return recipeList.map((recipe) => {
    return {
      params: {
        id: String(recipe.id),
      },
    };
  });
};

export const getRecipeData = async (id: string) => {
  const recipeList = await getAllRecipes()
  const data = recipeList.find((recipe) => String(recipe.id) === id);

  return {
    ...data,
  };
};
