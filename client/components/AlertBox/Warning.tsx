import AlertBox from "./MessageBox";

const Warning = ({ mainMessage, subMessage, onClick }) => {
  return (
    <div
      className="bg-orange-100 border border-orange-400 text-orange-700 px-4 py-3 rounded relative my-4"
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

export default Warning;
