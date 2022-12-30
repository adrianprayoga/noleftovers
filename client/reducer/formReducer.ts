const defaultIngredient = { name: "", amount: "", measure: "1" };
const defaultStep = { text: "" };

export const INITIAL_STATE = {
  name: "",
  description: "",
  ingredients: [{ ...defaultIngredient }],
  steps: [{ ...defaultStep }],
  author: "",
};

export const ACTION_TYPES = {
  UPDATE_TEXT: "UPDATE_TEXT",
  DELETE_ITEM: "DELETE_ITEM",
  LIST_UPDATE: "LIST_UPDATE",
  ADD_TO_LIST: "ADD_TO_LIST",
};

export const formReducer = (state, action) => {
  switch (action.type) {
    case ACTION_TYPES.UPDATE_TEXT:
      return {
        ...state,
        [action.payload.name]: action.payload.value,
      };
    case ACTION_TYPES.DELETE_ITEM:
      const pos = action.payload.position;
      return {
        ...state,
        [action.payload.name]: state[action.payload.name].filter(
          (_: any, i: number) => i !== pos
        ),
      };
    case ACTION_TYPES.LIST_UPDATE:
      const { name, position, value, prop } = action.payload;
      let newState = { ...state };
      newState[name][position][prop] = value;
      return newState;
    case ACTION_TYPES.ADD_TO_LIST:
      const def =
        action.payload.name === "ingredients" ? defaultIngredient : defaultStep;

      return {
        ...state,
        [action.payload.name]: [...state[action.payload.name], { ...def }],
      };
    default:
      return state;
  }
};
