import { useState, useContext, useEffect, useReducer } from "react";
import Layout from "../../components/Layout";
import Alert from "../../components/Alert";
import { CheckCircleIcon, ExclamationIcon } from "@heroicons/react/outline";
import { UserContext } from "../../hooks/userContext";
import clsx from "clsx";
import {
  ACTION_TYPES,
  userReducer,
  INITIAL_STATE,
} from "../../reducer/userReducer";
import { updateUser } from "../../lib/users";
import Success from "../../components/Success";

const Edit = () => {
  const userContext = useContext(UserContext);
  const [formState, dispatch] = useReducer(userReducer, INITIAL_STATE);
  const [validationError, setValidationError] = useState({});
  const [successMessage, setSuccessMessage] = useState(false);

  const handleChange = (event) => {
    dispatch({
      type: ACTION_TYPES.UPDATE_TEXT,
      payload: { name: event.target.name, value: event.target.value },
    });
  };

  const handleRemoveError = () => {
    let newValidationError = { ...validationError };
    delete newValidationError["overall"];
    setValidationError(newValidationError);
  };

  useEffect(() => {
    dispatch({
      type: ACTION_TYPES.UPDATE_TEXT,
      payload: { name: "name", value: userContext.user?.full_name },
    });
  }, [userContext]);

  const handleRecipeCreation = async () => {
    let newValidationError = {};

    if (!formState.name) {
      newValidationError["name"] = "this field is required";
    }

    if (Object.keys(newValidationError).length > 0) {
      setValidationError(newValidationError);
      return;
    }

    const response = await updateUser(formState["name"]);
    if (!response.error) {
      setSuccessMessage(true);
    } else {
      newValidationError["overall"] =
        "There seems to be an issue saving your user details. Please try again.";
    }

    setValidationError(newValidationError);
  };

  return (
    <Layout home={false} title="">
      <>
        {!userContext.authenticated ? (
          <h1 className="text-xl flex justify-center">
            You are currently not logged in. Please login to update your user
            details
          </h1>
        ) : (
          <div>
            <div className="md:grid md:grid-cols-1 md:gap-6">
              <div className="md:col-span-1">
                <div className="px-4 sm:px-0">
                  <h1 className="text-2xl font-medium leading-6 text-gray-900">
                    User Information
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
              {successMessage && (
                <Success
                  mainMessage="Success!"
                  subMessage={"Your details have been successfully saved"}
                  onClick={() => setSuccessMessage(false)}
                />
              )}
              <div className="mt-5 md:mt-0">
                <form action="#" method="POST">
                  <div className="shadow sm:rounded-md sm:overflow-hidden">
                    <div className="px-4 py-5 bg-white space-y-6 sm:p-6">
                      <div className="grid grid-cols-3 gap-6">
                        <div className="col-span-3">
                          <InputLabel label="Full Name" />
                          <div className="mt-1 flex shadow-sm">
                            <input
                              type="text"
                              name="name"
                              id="name"
                              className={clsx(
                                "focus:ring-indigo-500 focus:border-indigo-500 block w-full rounded-md sm:text-sm border-gray-300",
                                validationError["name"] && "border-red-500"
                              )}
                              placeholder="Full Name"
                              value={formState["name"] || ""}
                              onChange={handleChange}
                            />
                          </div>
                          <Error error={validationError["name"]} />
                        </div>
                      </div>
                      <div className="grid grid-cols-3 gap-6">
                        <div className="col-span-3">
                          <InputLabel label="Email" />
                          <div className="mt-1 flex shadow-sm">
                            <input
                              type="text"
                              name="email"
                              id="email"
                              className="block w-full rounded-md sm:text-sm border-gray-300 bg-gray-100"
                              placeholder="email"
                              value={userContext.user.email}
                              disabled
                            />
                          </div>
                        </div>
                      </div>
                    </div>

                    <div className="px-4 py-3 bg-gray-50 text-right sm:px-6">
                      <button
                        type="button"
                        className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        onClick={handleRecipeCreation}
                      >
                        Save
                      </button>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        )}
      </>
    </Layout>
  );
};

export default Edit;

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
