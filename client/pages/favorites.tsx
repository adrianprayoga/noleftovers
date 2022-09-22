import Head from "next/head";
import Layout, { siteTitle } from "../components/Layout";
import utilStyles from "../styles/utils.module.css";
import { getAllRecipes, getFavorites } from "../lib/recipes";
import RecipeGrid from "../components/RecipeGrid";
import Link from "next/link";
import { useContext, useEffect, useState } from "react";
import axios from "axios";
import useFavorites from "../hooks/favorites-hooks";
import { UserContext } from "../hooks/userContext";

axios.defaults.baseURL = `${process.env.NEXT_PUBLIC_BACKEND_HOST}`;
axios.defaults.withCredentials = true;

export default function Favorites({ allReciplesData }) {
  
  const userContext = useContext(UserContext);
  const { favorites, handleAddFavorite, handleRemoveFavorite } = useFavorites(userContext);

  return (
    <Layout home={false} title="Here's Your Favorite!">
      <Head>
        <title>{siteTitle}</title>
      </Head>

      {!userContext.authenticated ? (
        <h1 className="text-xl flex justify-center">
          You are currently not logged in. Please login to access your favorites
          recipe
        </h1>
      ) : (
        <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
          <RecipeGrid
            recipeList={allReciplesData.filter(
              (recipe) => favorites.indexOf(recipe.id) !== -1
            )}
            favorites={favorites}
            handleRemoveFavorite={handleRemoveFavorite}
          />
        </section>
      )}
    </Layout>
  );
}

export async function getStaticProps() {
  const allReciplesData = await getAllRecipes();

  return {
    props: { allReciplesData },
  };
}
