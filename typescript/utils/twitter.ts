import { TwitterApi } from "twitter-api-v2";
import { GeneratePath } from "./utils.js";

export default class Twitter {
  API_KEY_SECRET = "";
  API_KEY = "";
  TOKEN = {
    key: "",
    secret: "",
  };

  constructor() {}

  UpdateTwitterProfile = async (pseudo: string) => {
    const url = GeneratePath(pseudo);

    const client = new TwitterApi({
      appKey: this.API_KEY,
      appSecret: this.API_KEY_SECRET,
      accessToken: this.TOKEN.key,
      accessSecret: this.TOKEN.secret,
    });

    return await client.v1.updateAccountProfile({ name: pseudo });
  };
}
