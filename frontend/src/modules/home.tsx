import { Button } from "@fluentui/react";
import useHomeLogic, { IUseHomeLogic } from "./home.logic";

const Home = () => {
  const {
    actions: { setCounter, callbackFunc },
    state: { counter, myValue },
  }: IUseHomeLogic = useHomeLogic();

  return (
    <div>
      <Button
        onClick={() => {
          setCounter((c: number) => c + 1);
          callbackFunc(3);
        }}
      >
        Hello World
      </Button>
    </div>
  );
};
export default Home;
