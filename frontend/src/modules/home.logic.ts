import { useState, useEffect, useCallback, useMemo, useRef } from "react";

export interface IUseHomeLogic {
  actions: {
    callbackFunc: any;
    setCounter: any;
  };
  state: {
    counter: number;
    myValue: number;
  };
}
const useHomeLogic = () => {
  const [counter, setCounter] = useState(0);
  const myRef = useRef(0);

  useEffect(() => {
    setCounter(-1);

    myRef.current = 1;

    return () => {
      //
      setCounter(0);
    };
  }, []);

  const myValue = useMemo(() => {
    return counter * 10;
  }, [counter]);

  const callbackFunc = useCallback(
    (myparam: number) => {
      alert(myparam * counter);
    },
    [counter]
  );

  return {
    actions: {
      callbackFunc,
      setCounter,
    },
    state: {
      myValue,
      counter,
    },
  };
};
export default useHomeLogic;
