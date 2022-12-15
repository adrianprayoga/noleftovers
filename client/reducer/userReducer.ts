export const INITIAL_STATE = {
  name: ""
};

export const ACTION_TYPES = {
  UPDATE_TEXT: "UPDATE_TEXT",
};

export const userReducer = (state, action) => {
  switch (action.type) {
    case ACTION_TYPES.UPDATE_TEXT:
      return {
        ...state,
        [action.payload.name]: action.payload.value,
      };
    default:
      return state;
  }
};
