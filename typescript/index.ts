import Twitter from "./utils/twitter.js";
import { AzureFunction, Context } from "@azure/functions";
import { DaysBeforeCanada, GeneratePseudo } from "./utils/utils.js";

const timerTrigger: AzureFunction = async function (
  context: Context
): Promise<void> {
  context.log("Start script ts\n");

  const twitter = new Twitter();

  const days = DaysBeforeCanada();

  const pseudo = GeneratePseudo(days);

  context.log("Pseudo generate : %s\n", pseudo);

  try {
    twitter.UpdateTwitterProfile(pseudo);
    context.log("Script done");
  } catch (error) {
    context.log("Error during update profile : " + error);
  }
};

export default timerTrigger;
