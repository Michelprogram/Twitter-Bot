const travel = new Date("2025-01-01")
const url = "https://api.twitter.com/1.1/account/update_profile.json?name="
const canadianFlag = "🇨🇦"
const pseudo = "Bivouac - "

export const DaysBeforeCanada = () => {
	const now = new Date(Date.now())

    return Math.ceil((travel.getTime() - now.getTime()) / (1000 * 60 * 60 * 24));
}

export const GeneratePseudo = (days:number) :string =>{

    return `${pseudo}${days}J ${canadianFlag}`
}

export const GeneratePath = (pseudo:string) : string => {
	return url + pseudo
}
