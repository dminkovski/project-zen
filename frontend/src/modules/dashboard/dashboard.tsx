import { useEffect, useState } from "react";
import { Text } from "@fluentui/react";
import { IEmail } from "../../model/interfaces";

const Dashboard = () => {
  const [emails, setEmails] = useState([] as Array<IEmail>);

  const getEmails = async () => {
    const respondedEmails = await fetch(
      "https://project-zen.azurewebsites.net/project-zen/emails"
    );
    setEmails(respondedEmails as unknown as Array<IEmail>);
  };

  useEffect(() => {
    getEmails();
  }, []);

  return (
    <div>
      <Text
        as="h1"
        style={{
          fontSize: 38,
        }}
      >
        Project-Zen Dashboard
      </Text>
      {emails?.map((mail: IEmail) => (
        <div>
          <h4>{mail.Subject}</h4>
          <p>{mail.Body}</p>
        </div>
      ))}
    </div>
  );
};
export default Dashboard;
