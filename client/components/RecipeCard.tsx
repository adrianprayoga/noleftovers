import Link from "next/link";
import Image from "next/image";

interface RecipeHighlight {
  id: string;
  name: string;
  description: string;
}

const RecipeCard = (props: RecipeHighlight) => {
  return (
    <div className="justify-self-stretch flex flex-col align-top bg-white rounded-lg border shadow-md md:flex-row hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700">
      <img
        className="h-96 rounded-t-lg md:h-auto md:w-48 md:rounded-none md:rounded-l-lg"
        src={`/images/${props.id}.jpeg`}
        alt=""
        width="100%"
        height="100%"
      />
      <Link href={`/recipe/${props.id}`}>
        <div className="flex flex-col justify-between p-4 leading-normal">
          <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
            <a>{props.name}</a>
          </h5>
          <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
            {props.description}
          </p>
          <p>
            {`Tags: `}
            {["chicken", "burger", "pizza"].map((item) => {
              return (
                <Link href={`/${item}`}>
                  <a>{item + ", "}</a>
                </Link>
              );
            })}
          </p>
        </div>
      </Link>
    </div>
  );
};

export default RecipeCard;
