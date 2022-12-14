import { useState } from "react";

const SearchBar = (props) => {
  const [input, setInput] = useState("");

  return (
    <form
      onSubmit={e => {
        props.onClick(input);
        e.preventDefault()
        return false;
      }}
    >
      <label className="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-gray-300">
        Search
      </label>
      <div className="relative">
        <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
          <svg
            aria-hidden="true"
            className="w-5 h-5 text-gray-500 dark:text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
        </div>
        <input
          type="search"
          id="default-search"
          className="block p-2 pl-10 py-4 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
          placeholder="Search for recipes"
          value={input}
          onChange={(event) => setInput(event.target.value)}
        ></input>
        <button
          type="submit"
          className="text-white absolute right-0 bottom-0 border bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-r-lg text-sm px-4 py-4"
        >
          Search
        </button>
      </div>
    </form>
  );
};

export default SearchBar;
