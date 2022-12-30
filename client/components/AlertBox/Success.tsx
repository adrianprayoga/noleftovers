import AlertBox from "./MessageBox";

const Success = ({ mainMessage, subMessage, onClick }) => {
  return (
    <div
      className="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative my-4"
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

export default Success;
