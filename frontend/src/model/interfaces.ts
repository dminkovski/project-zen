export interface IEmail {
  From: string;
  Date: string;
  Subject: string;
  Body: string;
  Images?: string[];
  Links?: string[];
  Summary?: string;
}
