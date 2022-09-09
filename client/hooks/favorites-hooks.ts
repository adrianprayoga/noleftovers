import { useEffect, useState } from "react";
import { addFavorites, getFavorites, removeFavorites } from "../lib/recipes";

const useFavorites = () => {
  const [favorites, setFavorites] = useState([]);

  useEffect(() => {
    const getF = async () => {
      try {
        const res = await getFavorites();
        setFavorites(res);
        console.log(res);
      } catch (err) {
        console.error("call error", err);
      }
    };

    getF();
  }, []);

  const handleAddFavorite = async (recipe_id) => {
    try {
      await addFavorites(recipe_id);
      setFavorites(favorites.concat(recipe_id));
    } catch (error) {
      console.error(error);
    }
  };

  const handleRemoveFavorite = async (recipe_id) => {
    try {
      await removeFavorites(recipe_id);
      console.log(favorites.filter((f) => f !== recipe_id));
      setFavorites(favorites.filter((f) => f !== recipe_id));
    } catch (error) {
      console.error(error);
    }
  };

  return {
    favorites,
    handleAddFavorite,
    handleRemoveFavorite,
  };
};

export default useFavorites;
