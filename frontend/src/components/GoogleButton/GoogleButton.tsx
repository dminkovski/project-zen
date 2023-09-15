import { Button, ButtonProps } from "@mantine/core";
import { GoogleIcon } from "./GoogleIcon";
const GoogleButton = (
  props: ButtonProps & React.ComponentPropsWithoutRef<"a">
) => {
  return (
    <Button
      component="a"
      leftIcon={<GoogleIcon />}
      variant="default"
      color="gray"
      {...props}
    />
  );
};
export default GoogleButton;
