import { useEffect, useRef, useState } from "react";
import { Text } from "@fluentui/react";
import { IEmail } from "src/model/interfaces";
import { Spinner } from "@fluentui/react/lib/Spinner";

const Dashboard = () => {
  const [emails, setEmails] = useState([] as Array<IEmail>);
  const [loading, setLoading] = useState(false);

  const getEmails = async () => {
    if (!loading) {
      setLoading(true);
      try {
        const response = await fetch(
          "https://project-zen.azurewebsites.net/project-zen/emails",
          {
            method: "GET",
          }
        );
        const data = await response.json();
        if (data && Array.isArray(data) && data.length > 0) {
          const promises = [];
          for (let i = 0; i < data.length; i++) {
            const item = data[i];
            const promis = fetch("https://projectzen.azurewebsites.net", {
              method: "POST",
              body: JSON.stringify({
                message: (item as IEmail).Body,
              }),
              headers: { "Content-Type": "text/html; charset=utf-8" },
            });
            promises.push(promis);
          }
          const results = await Promise.all(promises);
          console.log(results);
          for (let j = 0; j < results.length; j++) {
            data[j].Summary = await results[j].text();
          }

          console.log(data);

          setEmails(data as unknown as Array<IEmail>);
        }
      } catch (error) {
        console.error(JSON.stringify(error));
      } finally {
        setLoading(false);
      }
    }
  };

  useEffect(() => {
    if (emails && Array.isArray(emails)) {
      setLoading(false);
    }
  }, [emails]);

  useEffect(() => {
    getEmails();
  }, []);

  return (
    <div style={{ background: "#f8f8f8" }}>
      <Text
        as="h1"
        style={{
          fontSize: 38,
        }}
      >
        Project-Zen Dashboard
      </Text>
      {loading && (
        <div>
          <Spinner
            label="Seriously, I am loading..."
            ariaLive="assertive"
            labelPosition="top"
          />
        </div>
      )}
      {!loading && emails?.length <= 0 && (
        <div>
          <p>No unread emails</p>
        </div>
      )}
      {emails &&
        Array.isArray(emails) &&
        emails?.map((mail: IEmail, index: number) => (
          <div
            key={index}
            style={{
              marginBottom: 25,
              background: "#fff",
              borderRadius: 15,
              padding: 15,
            }}
          >
            <h4>{mail.Subject}</h4>
            <div dangerouslySetInnerHTML={{ __html: mail.Body }} />
            <hr />
            <span style={{ color: "#666" }}>Summary:</span>
            <div dangerouslySetInnerHTML={{ __html: mail.Summary }} />
          </div>
        ))}
    </div>
  );
};
export default Dashboard;
