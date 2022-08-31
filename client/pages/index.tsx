import Head from "next/head";
import Layout, { siteTitle } from "../components/Layout";
import utilStyles from "../styles/utils.module.css";
import { getAllRecipes } from "../lib/recipes";
import RecipeGrid from "../components/RecipeGrid";
import Link from "next/link";
import { useEffect, useState } from "react";
import axios from "axios";

axios.defaults.baseURL = `${process.env.NEXT_PUBLIC_BACKEND_HOST}`;
axios.defaults.withCredentials = true;

export default function Home({ allReciplesData }) {
  const [appState, setAppState] = useState({
    user: {},
    error: null,
    authenticated: false,
  });

  // useEffect(() => {
  //   const getUser = async () => {
  //     try {
  //       const res = await axios.get(`/auth/success`);
  //       if (res.status == 200) {
  //         setAppState({
  //           user: res.data.user,
  //           authenticated: true,
  //           error: null,
  //         });
  //       }
  //       console.log(res);
  //     } catch (err) {
  //       console.error("call error", err);
  //       setAppState({
  //         user: {},
  //         authenticated: false,
  //         error: "user is not authenticated",
  //       });
  //     }
  //   };

  //   getUser();
  // }, []);

  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <RecipeGrid recipeList={allReciplesData} />
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
