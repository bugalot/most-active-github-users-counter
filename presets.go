package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "austria":PresetLocations{"austria", "%C3%B6sterreich", "vienna", "wien", "linz", "salzburg", "graz", "innsbruck", "klagenfurt", "wels", "dornbirn"},
  "finland":PresetLocations{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku"},
  "sweden":PresetLocations{"sweden", "sverige", "stockholm", "malm%C3%B6", "uppsala", "g%C3%B6teborg", "gothenburg"},
  "norway":PresetLocations{"norway", "norge", "oslo", "bergen"},
  "germany":PresetLocations{"germany", "deutschland", "berlin", "frankfurt", "munich", "m%C3%BCnchen", "hamburg", "cologne", "k%C3%B6ln"},
  "netherlands":PresetLocations{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland"},
  "ukraine":PresetLocations{"ukraine", "kiev", "kyiv", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
  "japan":PresetLocations{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
  "russia":PresetLocations{"russia", "moscow", "saint%2Bpetersburg", "novosibirsk", "yekaterinburg", "nizhny%2Bnovgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
  "estonia":PresetLocations{"estonia", "eesti", "tallinn", "tartu", "narva", "p%C3%A4rnu"},
  "denmark":PresetLocations{"denmark", "danmark", "copenhagen","aarhus","odense","aalborg"},
  "portugal":PresetLocations{"portugal", "lisbon", "lisboa","braga","porto","aveiro","coimbra","funchal", "madeira"},
  "france":PresetLocations{"france","paris","marseille","lyon","toulouse","nice","nantes","strasbourg","montpellier","bordeaux","lille","rennes","reims"},
  "spain":PresetLocations{"spain","espa%C3%B1a","madrid","barcelona","valencia","seville","sevilla","zaragoza","malaga","murcia","palma","bilbao","alicante","cordoba"},
  "italy":PresetLocations{"italy","italia","rome","roma","milan","naples","napoli","turin","torino","palermo","genoa","genova","bologna","florence","firenze","bari","catania","venice","verona"},
  "uk": PresetLocations{"uk","london","birmingham","leeds","glasgow","sheffield","bradford","manchester","edinburgh","liverpool","bristol","cardiff","belfast","leicester","wakefield","coventry","nottingham","newcastle"},
  "croatia": PresetLocations{"croatia","hrvatska","zagreb","split","rijeka","osijek","zadar","pula"},
  "worldwide": PresetLocations{},
  "china": PresetLocations{"china", "%E4%B8%AD%E5%9B%BD", "guangzhou", "shanghai", "beijing", "hangzhou", "hong%2Bkong"},
  "india": PresetLocations{"india", "mumbai", "delhi", "bangalore", "hyderabad", "ahmedabad", "chennai", "kolkata", "jaipur"},
  "indonesia": PresetLocations{"indonesia", "jakarta", "surabaya", "bandung", "medan", "bekasi", "semarang", "tangerang", "depok", "makassar", "palembang"},
  "pakistan": PresetLocations{"pakistan", "karachi", "lahore", "faisalabad", "rawalpindi", "peshawar", "islamabad"},
  "brazil": PresetLocations{"brazil", "brasil", "s%C3%A3o%20paulo", "bras%C3%ADlia", "salvador", "fortaleza", "belo%20horizonte", "manaus", "curitiba", "recife", "porto%20alegre"},
  "nigeria": PresetLocations{"nigeria", "lagos", "kano", "ibadan", "benin%20city", "port%20harcourt", "jos", "ilorin"},
  "bangladesh": PresetLocations{"bangladesh", "dhaka", "chittagong", "khulna", "rajshahi", "barisal", "sylhet", "rangpur", "comilla", "gazipur"},
  "mexico": PresetLocations{"mexico", "mexico%20city", "guadalajara", "puebla", "tijuana", "mexicali"},
  "philippines": PresetLocations{"philippines", "pilipinas", "quezon", "manila", "davao", "caloocan", "cebu", "zamboanga", "bohol", "pasig", "bacolod", "makati", "baguio"}}

func Preset(name string) []string {
  return PRESETS[name]
}
