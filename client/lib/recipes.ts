import axios from "axios";
import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();

export const getAllRecipes = async (): Promise<any[]> => {
  try {
    const response = await axios.get(`${process.env.HOST}/recipe`);

    return response.data;
  } catch (err) {
    console.error("error in getting recipes");

    return [];
  }
};

export const getAllRecipeIds = async () => {
  const recipeList = await getAllRecipes();

  return recipeList.map((recipe) => {
    return {
      params: {
        id: String(recipe.id),
      },
    };
  });
};

export const getRecipeData = async (id: string) => {
  try {
    const response = await axios.get(`${process.env.HOST}/recipe/${id}`);

    return response.data;
  } catch (err) {
    console.error("error in getting recipes");
    return {};
  }
};

export interface createRecipeEntry {
  name: string;
  description: string;
  ingredients: createIngredientsEntry[];
  steps: createStepsEntry[];
}

export interface createIngredientsEntry {
  name: string;
  amount: string;
  measure: string;
}

export interface createStepsEntry {
  text: string;
}

export const createRecipe = async (recipe: createRecipeEntry, file: string) => {
  let data = {
    name: recipe.name,
    description: recipe.description,
    ingredients: [],
    steps: [],
  };

  data.ingredients = recipe.ingredients
    .filter((i) => i.name)
    .map((ingredient) => {
      let temp = {};
      Object.keys(ingredient).forEach((k) => {
        if (ingredient[k]) {
          temp[k] = ingredient[k];
        }
      });
      return temp;
    });
  data.steps = recipe.steps.filter((i) => i.text);

  try {
    const formData = new FormData();
    formData.append("image", file);
    formData.append("recipe", JSON.stringify(data));

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_HOST}/recipe`,
      formData
    );
    return response.data;
  } catch (e) {
    console.error(e);
    return {
      error: true,
      status: e.response.status,
      errors: e.response.data,
      message: e.message,
    };
  }
};

export const searchRecipes = async (keys: string): Promise<any[]> => {
  try {
    const response = await axios.get(`${process.env.NEXT_PUBLIC_BACKEND_HOST}/recipe?search=${keys}`);

    return response.data;
  } catch (err) {
    console.error("error in getting recipes");

    return [];
  }
};

export const getMeasures = async (): Promise<any[]> => {
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_BACKEND_HOST}/measures`
    );

    return response.data;
  } catch (err) {
    console.error("error in getting measures", err);

    return [];
  }
};

export const getFavorites = async (): Promise<any[]> => {
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_BACKEND_HOST}/favorites`
    );
    return response.data;
  } catch (err) {
    console.error("error in getting favorites");
    return [];
  }
};

export const addFavorites = async (recipe_id: number): Promise<any> => {
  return axios.post(`${process.env.NEXT_PUBLIC_BACKEND_HOST}/favorites`, {
    recipe_id: recipe_id,
  });
};

export const removeFavorites = async (recipe_id: number): Promise<any> => {
  return axios.delete(
    `${process.env.NEXT_PUBLIC_BACKEND_HOST}/favorites/${recipe_id}`
  );
};
