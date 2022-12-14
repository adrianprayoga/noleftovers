import {
  useState,
  useCallback,
  useEffect,
  useReducer,
  useContext,
} from "react";
import Layout from "../../components/Layout";
import Alert from "../../components/AlertBox/Alert";
import { TrashIcon, ExclamationIcon, PlusIcon } from "@heroicons/react/outline";
import TextareaAutosize from "react-textarea-autosize";
import { createRecipe, getMeasures } from "../../lib/recipes";
import clsx from "clsx";
import ReactCrop, { Crop, PercentCrop } from "react-image-crop";
import "react-image-crop/dist/ReactCrop.css";
import {
  ACTION_TYPES,
  formReducer,
  INITIAL_STATE,
} from "../../reducer/formReducer";
import Warning from "../../components/AlertBox/Warning";
import { cropImage, dataURLtoFile, Rectangle } from "../../lib/images";
import { UserContext } from "../../hooks/userContext";

const Create = () => {
  const userContext = useContext(UserContext);

  console.log(userContext);

  const [formState, dispatch] = useReducer(formReducer, INITIAL_STATE);
  const [measures, setMeasures] = useState([]);
  const [file, setFile] = useState<File>(null);
  const [image, setImage] = useState("");
  const [imageSize, setImageSize] = useState(new Rectangle(0, 0));
  const [validationError, setValidationError] = useState({});
  const [crop, setCrop] = useState<PercentCrop>();

  function handleFileUpload(e) {
    try {
      setFile(e.target.files[0]);

      const imageUrl = URL.createObjectURL(e.target.files[0]);
      setImage(imageUrl);

      let img = new Image();
      img.src = imageUrl;
      img.onload = () => {
        setImageSize(new Rectangle(img.naturalWidth, img.naturalHeight));
      };
    } catch (e) {
      console.warn(e);
    }
  }

  function saveCropImage() {
    const { canvas, w, h } = cropImage(crop, image);

    const dataUrl = canvas.toDataURL(file.type, 1);

    setImage(dataUrl);
    setFile(dataURLtoFile(dataUrl, "cropped"));
    setImageSize(new Rectangle(w, h));

    setCrop(undefined);
  }

  useEffect(() => {
    const fetchMeasures = async () => {
      const data = await getMeasures();
      setMeasures(data);
    };

    fetchMeasures();
  }, []);

  const handleChange = (event) => {
    dispatch({
      type: ACTION_TYPES.UPDATE_TEXT,
      payload: { name: event.target.name, value: event.target.value },
    });
  };

  const handleItemDeletion = (name, position) => (event) => {
    dispatch({
      type: ACTION_TYPES.DELETE_ITEM,
      payload: { name: name, position: position },
    });
  };

  const handleListChange = (name: string, position: number) => (event) => {
    dispatch({
      type: ACTION_TYPES.LIST_UPDATE,
      payload: {
        name: name,
        position: position,
        prop: event.target.name,
        value: event.target.value,
      },
    });
  };

  const handleListAdd = (name: string) => (event) => {
    dispatch({
      type: ACTION_TYPES.ADD_TO_LIST,
      payload: {
        name: name,
      },
    });
  };

  const handleRemoveError = () => {
    let newValidationError = { ...validationError };
    delete newValidationError["overall"];
    setValidationError(newValidationError);
  };

  const handleRecipeCreation = async () => {
    let newValidationError = {};

    if (!formState.name) {
      newValidationError["name"] = "this field is required";
    }

    if (!formState.description) {
      newValidationError["description"] = "this field is required";
    }

    if (!file) {
      newValidationError["image"] = "please upload one image";
    }

    if (!imageSize.isRectangle) {
      newValidationError["image"] =
        "please crop your image to the correct aspect ratio";
    }

    if (formState.ingredients?.filter((e) => e.name).length === 0) {
      newValidationError["ingredients"] = "At least one ingredient is required";
    }

    if (formState.steps?.filter((e) => e.text).length === 0) {
      newValidationError["steps"] = "At least one recipe step is required";
    }

    if (Object.keys(newValidationError).length > 0) {
      setValidationError(newValidationError);
      return;
    }

    const author: string = userContext.user?.admin && formState['author'] ? formState['author'] : userContext.user.id
    const response = await createRecipe(formState, file, parseInt(author));
    if (!response.error) {
      window.location.href = `/recipe/${response.id}`;
    } else {
      newValidationError["overall"] =
        "There seems to be an issue creating your recipe. Please make sure that all inputs are correct.";
    }

    setValidationError(newValidationError);
  };

  return (
    <Layout home={false} title="">
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
            {validationError["overall"] && (
              <Alert
                mainMessage="Something went wrong!"
                subMessage={validationError["overall"]}
                onClick={handleRemoveError}
              />
            )}
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
                            className={clsx(
                              "focus:ring-indigo-500 focus:border-indigo-500 block w-full rounded-md sm:text-sm border-gray-300",
                              validationError["name"] && "border-red-500"
                            )}
                            placeholder="recipe name"
                            value={formState["name"]}
                            onChange={handleChange}
                          />
                        </div>
                        <Error error={validationError["name"]} />
                      </div>
                    </div>

                    {userContext.user?.admin && 
                      <div className="grid grid-cols-3 gap-6">
                        <div className="col-span-3 sm:col-span-2">
                          <InputLabel label="Author Id" />
                          <div className="mt-1 flex shadow-sm">
                            <input
                              type="number"
                              name="author"
                              id="author"
                              className={clsx(
                                "focus:ring-indigo-500 focus:border-indigo-500 block w-full rounded-md sm:text-sm border-gray-300"
                              )}
                              placeholder="author id"
                              value={formState["author"]}
                              onChange={handleChange}
                            />
                          </div>
                        </div>
                      </div>
                    }

                    <div>
                      <InputLabel label="Description" />
                      <div className="mt-1">
                        <TextareaAutosize
                          className={clsx(
                            "focus:ring-indigo-500 focus:border-indigo-500 block w-full rounded-md sm:text-sm border-gray-300",
                            validationError["description"] && "border-red-500"
                          )}
                          placeholder="brief description of the recipe"
                          value={formState["description"]}
                          onChange={handleChange}
                          minRows={5}
                          name="description"
                        />
                      </div>
                      <Error error={validationError["description"]} />
                    </div>

                    <div>
                      <InputLabel label="Picture" />
                      <div
                        className={clsx(
                          "mt-1 flex flex-col justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md",
                          validationError["image"] && "border-red-500"
                        )}
                      >
                        <div>
                          <div className="space-y-1 text-center pb-3">
                            {!file && (
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
                            )}

                            <div className="text-sm text-gray-600">
                              <label
                                htmlFor="file-upload"
                                className=" cursor-pointer bg-white rounded-md font-medium text-indigo-600 hover:text-indigo-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-indigo-500"
                              >
                                <span>{`${
                                  !file ? "Upload" : "Replace"
                                } your image`}</span>
                                <input
                                  id="file-upload"
                                  name="file-upload"
                                  type="file"
                                  className="sr-only"
                                  accept="image/x-png,image/jpg,image/jpeg"
                                  onChange={handleFileUpload}
                                />
                              </label>
                            </div>
                            <p className="text-xs text-gray-500">
                              PNG or JPG up to 10MB
                            </p>
                          </div>
                        </div>
                        <div>
                          {image && (
                            <div className="flex flex-col">
                              {!imageSize.isRectangle && (
                                <Warning
                                  mainMessage={""}
                                  subMessage={
                                    "Please use the resizing tool to crop your image into a square"
                                  }
                                  onClick={undefined}
                                />
                              )}

                              <div className="flex justify-center">
                                <button
                                  className="px-4 py-2 mb-2 rounded bg-indigo-100 text-indigo-600 hover:text-indigo-500 focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-indigo-500 flex-auto text-sm font-medium"
                                  type="button"
                                  onClick={saveCropImage}
                                >
                                  Crop Image
                                </button>
                              </div>
                              <ReactCrop
                                crop={crop}
                                onChange={(_, c) => setCrop(c)}
                                aspect={1}
                                className="justify-center"
                              >
                                <img
                                  src={image}
                                  alt="uploaded image"
                                  className="w-full"
                                />
                              </ReactCrop>
                            </div>
                          )}
                        </div>
                      </div>
                      <Error error={validationError["image"]} />
                    </div>

                    <div className="grid grid-cols-1 gap-1">
                      <InputLabel label="Ingredients List" />
                      {(formState["ingredients"] || []).map((item, i) => {
                        return (
                          <div className="mt-1 flex" key={i}>
                            <div className="mt-1 flex rounded-md w-full">
                              <span className="shadow-sm inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                                {i + 1}
                              </span>
                              <input
                                type="text"
                                className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-r-md sm:text-sm border-gray-300"
                                value={formState["ingredients"][i]["name"]}
                                name="name"
                                onChange={handleListChange("ingredients", i)}
                              />

                              <span className="inline-flex items-center text-sm text-gray-400 p-2">
                                amount
                              </span>
                              <input
                                type="text"
                                className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-1/4 rounded-md sm:text-sm border-gray-300"
                                placeholder="0"
                                value={formState["ingredients"][i]["amount"]}
                                name="amount"
                                onChange={handleListChange("ingredients", i)}
                              />
                            </div>

                            <select
                              autoComplete="measure"
                              className="ml-2 mt-1 block w-1/4 py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                              value={formState["ingredients"][i]["measure"]}
                              name="measure"
                              onChange={handleListChange("ingredients", i)}
                            >
                              {measures
                                .filter((measure) => measure.active)
                                .map((measure) => (
                                  <option key={measure.id} value={measure.id}>
                                    {measure.name}
                                  </option>
                                ))}
                            </select>

                            <button
                              type="button"
                              className="col-span-1 p-2"
                              name="test"
                              onClick={handleItemDeletion("ingredients", i)}
                            >
                              <TrashIcon className="h-6 w-6 text-gray-300" />
                            </button>
                          </div>
                        );
                      })}
                      <button
                        type="button"
                        className="col-span-1 p-2 m-1 mt-3"
                        name="test"
                        onClick={handleListAdd("ingredients")}
                      >
                        <div className="flex justify-center">
                          <PlusIcon className="h-5 w-5 text-indigo-400" />
                          <div className="text-sm ml-1 text-indigo-400">
                            Add New Ingredient
                          </div>
                        </div>
                      </button>
                      <Error error={validationError["ingredients"]} />
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
                                onChange={handleListChange("steps", i)}
                                name="text"
                                minRows={1}
                                maxRows={10}
                              />
                            </div>

                            <button
                              name="steps"
                              type="button"
                              className="col-span-1 p-2"
                              onClick={handleItemDeletion("steps", i)}
                            >
                              <TrashIcon className="h-6 w-6 text-gray-300" />
                            </button>
                          </div>
                        );
                      })}
                      <button
                        type="button"
                        className="col-span-1 p-2 m-1 mt-3"
                        name="test"
                        onClick={handleListAdd("steps")}
                      >
                        <div className="flex justify-center">
                          <PlusIcon className="h-5 w-5 text-indigo-400" />
                          <div className="text-sm ml-1 text-indigo-400">
                            Add New Steps
                          </div>
                        </div>
                      </button>
                      <Error error={validationError["steps"]} />
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

const Error = (props) => {
  const text = props.error ? props.error : "This is a required field";
  return (
    props.error && (
      <div className="mt-2 ml-1 flex">
        <ExclamationIcon className="h-4 w-4 text-red-500" />
        <div className="ml-1 text-xs text-red-500 align-middle">{text}</div>
      </div>
    )
  );
};
