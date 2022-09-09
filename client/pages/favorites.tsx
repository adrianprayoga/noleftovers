import Head from "next/head";
import Layout, { siteTitle } from "../components/Layout";
import utilStyles from "../styles/utils.module.css";
import { getAllRecipes, getFavorites } from "../lib/recipes";
import RecipeGrid from "../components/RecipeGrid";
import Link from "next/link";
import { useEffect, useState } from "react";
import axios from "axios";
import useFavorites from "../hooks/favorites-hooks";

axios.defaults.baseURL = `${process.env.NEXT_PUBLIC_BACKEND_HOST}`;
axios.defaults.withCredentials = true;

export default function Favorites({ allReciplesData, appState }) {
  const { favorites, handleAddFavorite, handleRemoveFavorite } = useFavorites();

  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>

      {/* {user: {…}, authenticated: true, error: null}
authenticated: true
error: null
user: {id: 0, email: 'aprayoga1994@gmail.com', password_hash: '', auth_method: '', oauth_id: '', …}
[[Prototype]]: Object */}

      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <RecipeGrid
          recipeList={allReciplesData.filter(
            (recipe) => favorites.indexOf(recipe.id) !== -1
          )}
          favorites={favorites}
          handleRemoveFavorite={handleRemoveFavorite}
        />
      </section>
      <div className="flex flex-col justify-between p-4 leading-normal">
        <button className="mb-2 text-lg font-bold tracking-tight text-gray-900">
          <Link href="/recipe/create">Create Recipe</Link>
        </button>
      </div>
    </Layout>
  );
}

export async function getStaticProps() {
  const allReciplesData = await getAllRecipes();

  return {
    props: { allReciplesData },
  };
}
