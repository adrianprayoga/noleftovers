import Head from "next/head";
import Layout, { siteTitle } from "../components/Layout";
import utilStyles from "../styles/utils.module.css";
import { getAllRecipes, searchRecipes } from "../lib/recipes";
import RecipeGrid from "../components/RecipeGrid";
import Link from "next/link";
import { useContext, useEffect, useState } from "react";
import axios from "axios";
import useFavorites from "../hooks/favorites-hooks";
import { UserContext } from "../hooks/userContext";
import SearchBar from "../components/SearchBar";

axios.defaults.baseURL = `${process.env.NEXT_PUBLIC_BACKEND_HOST}`;
axios.defaults.withCredentials = true;

export default function Home({ allReciplesData }) {
  const userContext = useContext(UserContext);
  const { favorites, handleAddFavorite, handleRemoveFavorite } =
    useFavorites(userContext);
  const [recipes, setRecipes] = useState(null);
  const [currentSearch, setCurrentSearch] = useState([]);

  const onClick = async (keys) => {
    const res = await searchRecipes(keys);
    setRecipes(res);
    setCurrentSearch(keys ? keys.split(",").map((s) => s.trim()) : []);
  };

  return (
    <Layout home title="Search for Your Recipe Here">
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <div className="mt-10">
        <SearchBar onClick={onClick} />
      </div>
      {currentSearch.length > 0 && (
        <div className="px-10, pt-2 text-gray-500">{`Currently showing search result for: ${currentSearch.join(
          ", "
        )}`}</div>
      )}
      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <RecipeGrid
          recipeList={recipes || allReciplesData}
          favorites={favorites}
          handleAddFavorite={handleAddFavorite}
          handleRemoveFavorite={handleRemoveFavorite}
        />
      </section>
    </Layout>
  );
}

export async function getStaticProps() {
  const allReciplesData = await getAllRecipes();

  return {
    props: { allReciplesData },
  };
}
