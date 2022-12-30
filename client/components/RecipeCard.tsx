import Link from "next/link";
import Image from "next/image";
import clsx from "clsx";
import { UserCircleIcon } from "@heroicons/react/outline";

interface NullString {
  String: string;
  Valid: boolean;
}

interface RecipeHighlight {
  id: string;
  name: string;
  description: string;
  author_name: NullString;
  imageLink: string;
  isFavorite: boolean;
  isDisabled: boolean;
  handleAddFavorite: (params: number) => number;
  handleRemoveFavorite: (params: number) => number;
}

const RecipeCard = (props: RecipeHighlight) => {
  return (
    // <div className="flex flex-col align-top bg-white rounded-lg border shadow-md md:flex-row hover:bg-gray-100">
    <div className="flex flex-col align-top bg-white rounded-lg border shadow-md md:flex-row hover:bg-gray-100">
      {props.imageLink && (
        <img
          className="hidden md:block rounded-t-lg md:h-auto md:w-48 md:rounded-none md:rounded-l-lg"
          src={`${process.env.NEXT_PUBLIC_BACKEND_HOST}/images/${props.imageLink}`}
          alt=""
        />
      )}

      <Link href={`/recipe/${props.id}`}>
        <div className="flex  flex-1 flex-col justify-start p-4 leading-normal cursor-pointer">
          <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900">
            <a>{props.name}</a>
          </h5>
          <p className=" mb-3 font-normal text-gray-700 text-base">
            {props.imageLink && (
              <img
                className="md:hidden h-36 w-36 rounded-lg float-left my-2 mr-4"
                src={`${process.env.NEXT_PUBLIC_BACKEND_HOST}/images/${props.imageLink}`}
                alt=""
              />
            )}
            {props.description?.length > 250
              ? props.description.substring(0, 250) + "..."
              : props.description}
          </p>

          <div className="flex text-gray-500">
            <UserCircleIcon className="h-5 w-5" />
            <div className="mx-2 text-sm">
              {`Author: ${
                props.author_name?.Valid ? props.author_name.String : "Unknown"
              }`}
            </div>
          </div>
        </div>
      </Link>

      <div
        className="m-4"
        onClick={() =>
          props.isFavorite
            ? props.handleRemoveFavorite(Number(props.id))
            : props.handleAddFavorite(Number(props.id))
        }
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className={clsx(
            "w-6 h-6",
            props.isDisabled && "stroke-gray-300",
            !props.isDisabled && "hover:fill-blue-400",
            props.isFavorite && "fill-blue-400"
          )}
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z"
          />
        </svg>
      </div>
    </div>
  );
};

export default RecipeCard;
