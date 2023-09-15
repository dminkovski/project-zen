import { useEffect, useState } from "react";
import { IEmail } from "../../model/interfaces";
import Footer from "../../components/Footer/Footer";
import { Loader, Text } from "@mantine/core";
import { Email } from "../../components/Email/Email";
import * as Config from "../../config/config";

const Dashboard = () => {
  const [emails, setEmails] = useState([] as Array<IEmail>);
  const [summary, setSummary] = useState("");
  const [loading, setLoading] = useState(false);

  const getEmails = async () => {
    if (!loading) {
      setLoading(true);
      try {
        const response = await fetch(Config.EMAILS_API, {
          method: "GET",
        });
        const data = await response.json();
        const mails = data.Mails;
        const summary = data.Summary;
        setSummary(summary);
        if (mails && Array.isArray(mails) && mails.length > 0) {
          setEmails(mails as unknown as Array<IEmail>);
        }
      } catch (error) {
        setSummary("");
        setEmails([]);
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
    <div style={{ position: "relative", paddingBottom: 100, minHeight: 600 }}>
      <img
        src="logo.png" // Add your logo image here
        style={{ width: 100, margin: "auto", marginBottom: "50px" }}
      />
      {loading && (
        <div>
          <Loader />
          <br />
          <Text
            variant="gradient"
            gradient={{ from: "indigo", to: "cyan", deg: 45 }}
            sx={{ fontFamily: "Greycliff CF, sans-serif" }}
            ta="center"
            fz="xl"
            fw={300}
          >
            Patience, I am loading...grab a coffee.
          </Text>
        </div>
      )}
      {summary && !loading && (
        <div>
          <Text
            size="xl"
            variant="gradient"
            gradient={{ from: "green", to: "black", deg: 45 }}
            mb={25}
          >
            Executive Summary
          </Text>
          <Text size="l" color="" align="justify">
            {summary}
          </Text>
        </div>
      )}

      {!loading && emails?.length <= 0 && (
        <div>
          <Text
            size="xl"
            variant="gradient"
            gradient={{ from: "green", to: "black", deg: 45 }}
            mb={25}
            mt={50}
          >
            No unread newsletters
          </Text>
        </div>
      )}
      {emails && Array.isArray(emails) && emails.length > 0 && (
        <div>
          <Text
            size="xl"
            mb={25}
            mt={50}
            variant="gradient"
            gradient={{ from: "green", to: "black", deg: 45 }}
          >
            Newsletters
          </Text>
          {emails?.map((mail: IEmail, index: number) => (
            <div key={index}>
              <Email
                title={mail.Subject}
                from={mail.From}
                body={mail.Body}
                date={mail.Date}
              />
            </div>
          ))}{" "}
        </div>
      )}
      <Footer />
    </div>
  );
};
export default Dashboard;
