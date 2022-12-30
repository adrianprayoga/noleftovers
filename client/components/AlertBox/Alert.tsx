import AlertBox from "./MessageBox";

const Alert = ({ mainMessage, subMessage, onClick }) => {
  return (
    <div
      className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative my-4"
      role="alert"
    >
      <AlertBox
        mainMessage={mainMessage}
        subMessage={subMessage}
        onClick={onClick}
      />
    </div>
  );
};

export default Alert;
