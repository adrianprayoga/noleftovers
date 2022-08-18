import { useState, useCallback, useEffect } from "react";
import Layout from "../../components/Layout";
import { TrashIcon } from "@heroicons/react/outline";
import TextareaAutosize from "react-textarea-autosize";
import { createRecipe, getMeasures } from "../../lib/recipes";

const Create = (props) => {
  const { recipeData } = props;
  const [formState, setFormState] = useState({
    name: "",
    description: "",
    ingredients: [{ name: "", amount: "", measure: "1" }],
    steps: [{ text: "" }],
  });
  const [measures, setMeasures] = useState([]);
  const [file, setFile] = useState("");

  function handleFileUpload(e) {
    console.log(e.target.files);
    setFile(e.target.files[0]);
  }

  useEffect(() => {
    const fetchMeasures = async () => {
      const data = await getMeasures();
      setMeasures(data);
    };

    fetchMeasures();
  }, []);

  const handleChange = (name) => (event) => {
    setFormState((state) => ({
      ...state,
      [name]: event.target.value,
    }));
  };

  const handleListChange =
    (name: string, prop: string, item_i: number) => (event) => {
      setFormState((state) => {
        let newState = { ...state };
        newState[name][item_i][prop] = event.target.value;

        const lastItem = newState[name][newState[name].length - 1];
        if (Object.values(lastItem).find((i) => i)) {
          const newDefault = { ...lastItem };
          Object.keys(newDefault).forEach((k) => (newDefault[k] = ""));

          newState[name] = newState[name].concat(newDefault);
        }

        return newState;
      });
    };

  const handleRecipeCreation = async () => {
    const response = await createRecipe(formState, file);
    console.log(response);
    if (!response.error) {
      window.location.href = `http://localhost:3000/recipe/${response.id}`;
    }
  };

  const handleItemDeletion = (name, item_i) => {
    setFormState((state) => {
      let newState = { ...state };
      newState[name] = newState[name].filter((_, i) => i !== item_i);

      return newState;
    });
  };

  return (
    <Layout home={false}>
      <>
        <div>
          <div className="md:grid md:grid-cols-1 md:gap-6">
            <div className="md:col-span-1">
              <div className="px-4 sm:px-0">
                <h1 className="text-2xl font-medium leading-6 text-gray-900">
                  Create New Recipe
                </h1>
                <p className="mt-1 text-sm text-gray-600"></p>
              </div>
            </div>
            <div className="mt-5 md:mt-0">
              <form action="#" method="POST">
                <div className="shadow sm:rounded-md sm:overflow-hidden">
                  <div className="px-4 py-5 bg-white space-y-6 sm:p-6">
                    <div className="grid grid-cols-3 gap-6">
                      <div className="col-span-3 sm:col-span-2">
                        <InputLabel label="Recipe Name" />
                        <div className="mt-1 flex shadow-sm">
                          <input
                            type="text"
                            name="name"
                            id="name"
                            className="focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-md sm:text-sm border-gray-300"
                            placeholder="recipe name"
                            value={formState["name"]}
                            onChange={handleChange("name")}
                          />
                        </div>
                      </div>
                    </div>

                    <div>
                      <InputLabel label="Description" />
                      <div className="mt-1">
                        <TextareaAutosize
                          className=" shadow-sm focus:ring-indigo-500 focus:border-indigo-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"
                          placeholder="brief description of the recipe"
                          value={formState["description"]}
                          onChange={handleChange("description")}
                          minRows={5}
                        />
                      </div>
                      {/* <p className="mt-2 text-sm text-gray-500">
                        Brief description for your profile. URLs are
                        hyperlinked.
                      </p> */}
                    </div>

                    <div>
                      <InputLabel label="Picture" />
                      <div className="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
                        <div className="space-y-1 text-center">
                          <svg
                            className="mx-auto h-12 w-12 text-gray-400"
                            stroke="currentColor"
                            fill="none"
                            viewBox="0 0 48 48"
                            aria-hidden="true"
                          >
                            <path
                              d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02"
                              strokeWidth={2}
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                          </svg>
                          <div className="flex text-sm text-gray-600">
                            <label
                              htmlFor="file-upload"
                              className="relative cursor-pointer bg-white rounded-md font-medium text-indigo-600 hover:text-indigo-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-indigo-500"
                            >
                              <span>Upload a file</span>
                              <input
                                id="file-upload"
                                name="file-upload"
                                type="file"
                                className="sr-only"
                                onChange={handleFileUpload}
                              />
                            </label>
                            {file && (
                                <img
                                  src={file}
                                  width="300px"
                                  height="300px"
                                  alt="selected image..."
                                />
                                
                            )}

                            <p className="pl-1">or drag and drop</p>
                          </div>
                          <p className="text-xs text-gray-500">
                            PNG or JPG up to 10MB
                          </p>
                        </div>
                      </div>
                    </div>

                    <div className="grid grid-cols-1 gap-1">
                      <InputLabel label="Ingredients List" />
                      {(formState["ingredients"] || []).map((item, i) => {
                        return (
                          <div className="mt-1 flex" key={i}>
                            <div className="mt-1 flex rounded-md shadow-sm w-full">
                              <span className="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                                {i + 1}
                              </span>
                              <input
                                type="text"
                                className="focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-r-md sm:text-sm border-gray-300"
                                value={formState["ingredients"][i]["name"]}
                                onChange={handleListChange(
                                  "ingredients",
                                  "name",
                                  i
                                )}
                              />

                              <span className="inline-flex items-center text-sm text-gray-400 p-2">
                                amount
                              </span>
                              <input
                                type="text"
                                className="focus:ring-indigo-500 focus:border-indigo-500 block w-1/4 rounded-md sm:text-sm border-gray-300"
                                placeholder="0"
                                value={formState["ingredients"][i]["amount"]}
                                onChange={handleListChange(
                                  "ingredients",
                                  "amount",
                                  i
                                )}
                              />
                            </div>

                            <select
                              autoComplete="measure"
                              className="ml-2 mt-1 block w-1/4 py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                              value={formState["ingredients"][i]["measure"]}
                              onChange={handleListChange(
                                "ingredients",
                                "measure",
                                i
                              )}
                            >
                              {measures
                                .filter((measure) => measure.active)
                                .map((measure) => (
                                  <option key={measure.id} value={measure.id}>
                                    {measure.name}
                                  </option>
                                ))}
                            </select>

                            {formState["ingredients"]?.length > 1 && (
                              <button
                                type="button"
                                className="col-span-1 p-2"
                                onClick={() =>
                                  handleItemDeletion("ingredients", i)
                                }
                              >
                                <TrashIcon className="h-6 w-6 text-gray-300" />
                              </button>
                            )}
                          </div>
                        );
                      })}
                    </div>

                    <div className="grid grid-cols-1 gap-1">
                      <InputLabel label="Steps" />
                      {(formState["steps"] || []).map((item, i) => {
                        return (
                          <div className="mt-1 flex" key={i}>
                            <div className="mt-1 flex rounded-md shadow-sm w-full">
                              <span className="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                                {i + 1}
                              </span>
                              <TextareaAutosize
                                className="focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-r-md sm:text-sm border-gray-300"
                                value={formState["steps"][i]["text"]}
                                onChange={handleListChange("steps", "text", i)}
                                minRows={1}
                                maxRows={10}
                              />
                            </div>

                            {formState["steps"]?.length > 1 && (
                              <button
                                type="button"
                                className="col-span-1 p-2"
                                onClick={() => handleItemDeletion("steps", i)}
                              >
                                <TrashIcon className="h-6 w-6 text-gray-300" />
                              </button>
                            )}
                          </div>
                        );
                      })}
                    </div>
                  </div>
                  <div className="px-4 py-3 bg-gray-50 text-right sm:px-6">
                    <button
                      type="button"
                      className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                      onClick={handleRecipeCreation}
                    >
                      Create Recipe
                    </button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </>
    </Layout>
  );
};

export default Create;

const InputLabel = ({ label }) => {
  return (
    <label htmlFor={label} className="block text-sm font-medium text-gray-700">
      {label}
    </label>
  );
};
