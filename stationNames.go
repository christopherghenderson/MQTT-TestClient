package mqtttestclient

var randomMessageType = []string{"analog", "digital"}

var randomDevicePrefix = []string{"pt", "dt", "tt", "mov", "ft", "vt"}

var randomStationPrefix = []string{"crate", "trade", "mouth", "flower", "pest", "coat", "jeans", "fold", "air", "minister", "base", "grandmother", "smash", "knife", "steam", "children", "shelf", "shock", "reward", "giraffe", "basin", "poison", "airplane", "cake", "sticks", "yard", "wrist", "vessel", "color", "pump", "drain", "farm", "box", "planes", "snakes", "curve", "story", "wheel", "crib", "rate", "cars", "carriage", "sky", "cows", "recess", "trip", "pollution", "icicle", "hate", "person", "seashore", "discovery", "spark", "title", "crack", "song", "zoo", "day", "destruction", "glove", "science", "guide", "error", "pets", "education", "books", "lamp", "lace", "language", "downtown", "furniture", "bike", "quiet", "tail", "screw", "push", "foot", "sleep", "fang", "church", "instrument", "history", "play", "end", "friend", "fish", "floor", "suit", "statement", "point", "weather", "sister", "girl", "distribution", "sack", "vacation", "grape", "stove", "afternoon", "drop", "book", "record", "ship", "self", "burst", "cannon", "plantation", "market", "zebra", "man", "linen", "income", "wax", "oranges", "skin", "twig", "flowers", "waves", "zinc", "quiver", "observation", "distance", "cause", "dime", "quilt", "voyage", "ground", "table", "rice", "toad", "throat", "step", "number", "woman", "school", "friends", "car", "rock", "stretch", "wrench", "roof", "popcorn", "toy", "actor", "muscle", "thing", "underwear", "silk", "skate", "clover", "bird", "duck", "rule", "way", "club", "meat", "system", "elbow", "giants", "pin", "knowledge", "cats", "lunch", "toothpaste", "bucket", "basketball", "frame", "cap", "cow", "money", "dirt", "railway", "paper", "verse", "power", "argument", "ball", "division", "bushes", "fireman", "queen", "string", "cobweb", "roll", "angle", "square", "weight", "change", "straw", "grandfather", "transport", "bead", "force", "rest", "field", "account", "kiss", "competition", "line", "wash", "health", "birthday", "circle", "coast", "gun", "mask", "snails", "connection", "powder", "rose", "pull", "houses", "teaching", "copper", "war", "experience", "wilderness", "question", "place", "sink", "store", "trees", "mountain", "boot", "scissors", "letter", "door", "lumber", "hand", "pet", "truck", "wound", "design", "year", "toothbrush", "tank", "earthquake", "badge", "mailbox", "writer", "wool", "back", "plastic", "rub", "vase", "page", "haircut", "tax", "office", "cherry", "camera", "vegetable", "knot", "minute", "stage", "cat", "spy", "government", "spring", "direction", "battle", "hot", "twist", "calculator", "increase", "ocean", "blood", "debt", "trail", "picture", "desk", "teeth", "yarn", "daughter", "head", "station", "quicksand", "party", "plant", "calendar", "trick", "rainstorm", "cattle", "creator", "finger", "spot", "drawer", "night", "bat", "pizzas", "cakes", "gate", "stranger", "receipt", "van", "current", "wealth", "bells", "drink", "smoke", "apparatus", "stew", "notebook", "comparison", "beginner", "crow", "plot", "request", "dust", "geese", "orange", "middle", "hydrant", "stamp", "tree", "stone", "guitar", "brother", "wire", "soap", "temper", "match", "tramp", "tent", "pancake", "jam", "rat", "silver", "belief", "mind", "position", "horn", "amount", "throne", "rain", "bee", "loaf", "regret", "writing", "shop", "test", "degree", "arithmetic", "cheese", "sponge", "egg", "waste", "summer", "bottle", "pan", "fairies", "wren", "insect", "sugar", "doctor", "mark", "finger", "island", "vest", "walk", "vein", "squirrel", "memory", "lake", "coil", "hat", "nerve", "cabbage", "celery", "cushion", "insurance", "credit", "needle", "women", "clam", "interest", "dad", "hour", "condition", "word", "stop", "run", "swim", "week", "snow", "tomatoes", "jellyfish", "cough", "purpose", "idea", "oil", "songs", "soda", "tin", "rifle", "ray", "men", "discussion", "donkey", "feeling", "plate", "agreement", "turn", "quartz", "board", "brake", "appliance", "brass", "approval", "development", "hall", "camp", "stem", "tray", "harbor", "heat", "meeting", "room", "channel", "start", "chickens", "birds", "boat", "horses", "rings", "activity", "payment", "toes", "curtain", "punishment", "arm", "texture", "monkey", "mine", "effect", "holiday", "seat", "arch", "tooth", "rabbits", "use", "bulb", "mass", "carpenter", "measure", "blow", "milk", "relation", "wood", "show", "pencil", "committee", "sand", "stream", "basket", "love", "stitch", "library", "porter", "army", "cook", "plane", "ear", "servant", "anger", "sleet", "laborer", "scale", "potato", "look", "nose", "canvas", "things", "cherries", "ducks", "quill", "invention", "crook", "hair", "food", "oatmeal", "creature", "cream", "flag", "side", "month", "behavior", "join", "eyes", "front", "turkey", "rod", "chin", "hook", "honey", "surprise", "attraction", "thunder", "bit", "grain", "hobbies", "loss", "bikes", "window", "cloth", "wave", "reason", "hands", "coach", "zephyr", "acoustics", "addition", "treatment", "business", "dress", "wing", "pig", "touch", "sneeze", "taste", "button", "glass", "lock", "noise", "root", "doll", "shape", "slope", "flesh", "plants", "believe", "star", "beds", "breath", "scene", "name", "country", "fruit", "quince", "babies", "sea", "sheep", "alarm", "grass", "veil", "cast", "oven", "locket", "bell", "hammer", "seed", "thumb", "range", "passenger", "thrill", "bubble", "collar", "jump", "butter", "cup", "blade", "hole", "birth", "lip", "pail", "dinner", "rake", "son", "offer", "baby", "flock", "spiders", "building", "expert", "crown", "advice", "rabbit", "produce", "view", "iron", "talk", "religion", "smile", "winter", "water", "sisters", "care", "ladybug", "afterthought", "bedroom", "scarf", "machine", "trucks", "flavor", "playground", "land", "can", "suggestion", "shirt", "dogs", "move", "yak", "profit", "cemetery", "trains", "ring", "mint", "snake", "event", "nest", "whip", "dog", "slip", "voice", "key", "apparel", "shame", "house", "cactus", "brick", "engine", "top", "bed", "toe", "mist", "rail", "toys", "scarecrow", "act", "fire", "protest", "zipper", "face", "existence", "aunt", "fall", "pen", "detail", "sidewalk", "skirt", "stick", "frog", "tub", "maid", "low", "laugh", "sign", "rhythm", "animal", "visitor", "shade", "impulse", "harmony", "neck", "stocking", "route", "wish", "train", "trouble", "pigs", "whistle", "slave", "beef", "branch", "partner", "tongue", "mitten", "scent", "team", "fog", "dock", "trousers", "coal", "quarter", "ice", "robin", "riddle", "worm", "swing", "liquid", "sock", "knee", "baseball", "tiger", "sort", "sweater", "aftermath", "kitty", "flame", "snail", "stomach", "pocket", "mom", "fly", "chess", "reaction", "volcano", "wine", "achiever", "art", "kick", "note", "fowl", "border", "caption", "authority", "territory", "group", "balance", "desire", "mice", "expansion", "yam", "hill", "town", "jewel", "juice", "theory", "dinosaurs", "level", "watch", "lettuce", "cub", "steel", "gold", "class", "thread", "uncle", "fact", "pot", "reading", "flight", "nation", "cart", "motion", "company", "death", "mother", "sail", "deer", "airport", "humor", "hospital", "frogs", "kittens", "need", "pies", "support", "cover", "secretary", "parcel", "spade", "pleasure", "boy", "pie", "value", "bag", "cable", "amusement", "road", "control", "pear", "respect", "grade", "river", "sense", "home", "eye", "grip", "limit", "fork", "cracker", "sun", "driving", "corn", "north", "sound", "ink", "chalk", "work", "berry", "adjustment", "bone", "representative", "shake", "plough", "bridge", "pickle", "growth", "wind", "umbrella", "spoon", "volleyball", "pipe", "metal", "nut", "morning", "kettle", "hope", "structure", "attack", "sofa", "crime", "edge", "card", "leg", "cent", "hose", "eggs", "moon", "wall", "sheet", "jelly", "tendency", "lunchroom", "bait", "example", "jar", "bath", "industry", "marble", "ticket", "yoke", "eggnog", "crowd", "legs", "fear", "salt", "space", "part", "unit", "digestion", "form", "crayon", "girls", "governor", "leather", "fuel", "judge", "bomb", "bear", "time", "size", "cellar", "ghost", "substance", "earth", "property", "smell", "society", "thought", "action", "friction", "street", "letters", "boundary", "dolls", "shoe", "chicken", "price", "advertisement", "exchange", "ants", "selection", "peace", "shoes", "jail", "soup", "meal", "bite", "chance", "order", "magic", "horse", "cave", "decision"}

var randomStationSuffix = []string{"crate", "meek", "waves", "grip", "victorious", "name", "eight", "appliance", "crawl", "rule", "productive", "judge", "scrub", "low", "utopian", "pot", "reading", "can", "mountainous", "trail", "young", "match", "handy", "snail", "boorish", "rub", "towering", "tax", "tiny", "plastic", "adamant", "bomb", "substance", "terrible", "thing", "things", "abortive", "hot", "racial", "property", "copper", "branch", "cloistered", "meeting", "different", "fragile", "sense", "order", "cracker", "receipt", "allow", "abusive", "bubble", "damaging", "wrong", "lopsided", "sticky", "elastic", "divergent", "super", "business", "suck", "unfasten", "pull", "ablaze", "sand", "spark", "equable", "claim", "giant", "arrange", "stamp", "unbiased", "humor", "ceaseless", "filthy", "needless", "ancient", "canvas", "acidic", "abiding", "incompetent", "sore", "debonair", "fasten", "unaccountable", "smiling", "flowers", "black-and-white", "paint", "cushion", "smooth", "connection", "interesting", "truck", "flavor", "silk", "water", "laugh", "communicate", "smart", "representative", "top", "uptight", "arrive", "afford", "hour", "malicious", "tawdry", "birds", "good", "pies", "public", "grade", "comfortable", "combative", "royal", "concentrate", "superficial", "float", "outrageous", "telephone", "slippery", "queen", "healthy", "wiggly", "matter", "talented", "annoying", "smile", "numberless", "reproduce", "relax", "panicky", "fireman", "relieved", "advise", "back", "berry", "hall", "bent", "undress", "frequent", "encouraging", "science", "double", "bite", "ink", "icicle", "measure", "vein", "oval", "kiss", "tasty", "reflective", "unite", "quizzical", "receive", "invention", "impossible", "pastoral", "trip", "rambunctious", "lonely", "dry", "blade", "list", "record", "pretty", "secretive", "needy", "run", "crow", "interest", "sock", "worm", "guarded", "launch", "team", "peck", "many", "men", "colossal", "love", "destruction", "turn", "naive", "driving", "order", "excited", "babies", "debt", "slimy", "lackadaisical", "material", "admit", "assorted", "start", "bikes", "well-to-do", "beam", "library", "overt", "tall", "cultured", "lick", "knit", "knot", "tame", "scent", "sparkle", "scattered", "ray", "balance", "argument", "birth", "walk", "vase", "obtain", "temporary", "impress", "descriptive", "hulking", "fluffy", "certain", "picture", "telling", "internal", "upset", "military", "excuse", "attach", "women", "cast", "clever", "panoramic", "muscle", "cherry", "letters", "wing", "direful", "average", "weight", "flag", "flap", "mom", "easy", "crate", "teaching", "even", "root", "unkempt", "incandescent", "substantial", "quiver", "cheerful", "loutish", "vagabond", "capable", "promise", "hospitable", "craven", "languid", "act", "spectacular", "calculating", "itchy", "heat", "lake", "irritate", "unequaled", "decisive", "spotted", "cobweb", "noxious", "face", "capricious", "mind", "legs", "overconfident", "exultant", "burst", "weary", "lacking", "enormous", "drunk", "believe", "fertile", "dock", "writing", "purring", "abashed", "wheel", "house", "succeed", "painstaking", "brake", "oceanic", "judicious", "tie", "alive", "cap", "absorbing", "wholesale", "waste", "motionless", "cagey", "person", "shelf", "godly", "value", "government", "cows", "well-off", "partner", "support", "sad", "longing", "milk", "bucket", "oven", "blind", "astonishing", "store", "dizzy", "decay", "remarkable", "squeak", "redundant", "deranged", "wink", "cut", "calculate", "standing", "knot", "quack", "unruly", "snails", "lucky", "pretend", "mint", "sign", "ignorant", "soup", "large", "houses", "rustic", "parallel", "feeling", "injure", "mug", "verse", "cover", "shivering", "acid", "terrify", "tranquil", "face", "coach", "reach", "kittens", "helpless", "volleyball", "dinner", "polite", "thankful", "crib", "ragged", "file", "pause", "scorch", "behavior", "jittery", "miscreant", "vast", "fallacious", "frightening", "squealing", "credit", "dog", "wanting", "chase", "guess", "develop", "pink", "camp", "long-term", "drop", "swing", "pumped", "arm", "shoes", "careless", "wistful", "butter", "yak", "enchanting", "spray", "straight", "thumb", "waggish", "fuzzy", "plain", "scribble", "absorbed", "live", "river", "medical", "cream", "worry", "protect", "wilderness", "crack", "month", "mice", "class", "serve", "dress", "witty", "show", "challenge", "cattle", "quill", "woman", "polish", "pick", "suppose", "frog", "print", "turkey", "prefer", "quirky", "receptive", "inconclusive", "snatch", "way", "abrasive", "replace", "hurt", "succinct", "graceful", "trite", "hallowed", "clear", "pigs", "gold", "accidental", "agree", "sedate", "dull", "warlike", "intelligent", "need", "string", "harm", "macho", "railway", "check", "deliver", "learned", "tremble", "amusement", "threatening", "inform", "unusual", "tub", "apathetic", "familiar", "cart", "structure", "slap", "owe", "unbecoming", "strengthen", "heartbreaking", "dare", "low", "smash", "mist", "honey", "street", "train", "swim", "company", "inquisitive", "tacky", "weak", "join", "aboard", "bang", "dependent", "minute", "bow", "frantic", "ice", "messy", "pine", "bite-sized", "lock", "kitty", "ethereal", "trip", "vacation", "force", "earn", "absurd", "exercise", "wise", "knee", "club", "historical", "highfalutin", "box", "spy", "duck", "unknown", "basketball", "auspicious", "adjoining", "balance", "north", "quilt", "hot", "ubiquitous", "unnatural", "yawn", "furtive", "preach", "children", "release", "coil", "moaning", "grape", "rainstorm", "abject", "repulsive", "measly", "disturbed", "crayon", "box", "drain", "crack", "save", "drawer", "pie", "peel", "bore", "finicky", "scratch", "ball", "sip", "hope", "greedy", "shirt", "bells", "metal", "tough", "detail", "miniature", "present", "jail", "plant", "retire", "tricky", "deafening", "grate", "deep", "veil", "wrap", "gleaming", "beg", "confuse", "marry", "preserve", "mess up", "idiotic", "self", "word", "gaping", "leg", "cabbage", "fruit", "grandfather", "quarrelsome", "delirious", "field", "ludicrous", "honorable", "special", "daffy", "rude", "addicted", "glossy", "neat", "groan", "alarm", "rabbit", "eager", "play", "elated", "stupid", "jar", "unlock", "spiky", "enchanted", "infamous", "wrathful", "complex", "glass", "sincere", "rub", "soap", "hellish", "influence", "whistle", "acoustic", "sturdy", "increase", "flock", "hand", "obsequious", "brash", "gruesome", "tire", "deserve", "insidious", "dust", "sponge", "week", "repair", "fear", "present", "porter", "rampant", "brawny", "white", "cable", "domineering", "faint", "fairies", "jam", "society", "prepare", "airplane", "quaint", "fold", "trot", "grumpy", "hospital", "country", "glistening", "fence", "riddle", "misty", "thoughtful", "hope", "rigid", "own", "distribution", "stormy", "charming", "erect", "aromatic", "permit", "note", "idea", "undesirable", "snobbish", "able", "beneficial", "pig", "marked", "like", "cure", "noise", "knowing", "rush", "wreck", "innocent", "floor", "sassy", "aware", "belief", "frogs", "hill", "crowd", "boundary", "reflect", "brake", "race", "pickle", "open", "hunt", "scrawny", "grass", "clam", "wealthy", "fluttering", "resonant", "psychedelic", "dazzling", "lip", "drab", "cloth", "unsuitable", "existence", "green", "next", "simplistic", "disastrous", "cheese", "shake", "probable", "burly", "weather", "interfere", "entertain", "drum", "shelter", "perpetual", "profit", "use", "frame", "laborer", "lighten", "produce", "confess", "wakeful", "questionable", "sun", "regret", "elfin", "flawless", "quince", "tasteless", "efficacious", "steer", "spoon", "venomous", "unwieldy", "rapid", "bored", "fanatical", "dangerous", "plant", "knotty", "protest", "empty", "soft", "clumsy", "cherries", "sound", "happen", "money", "big", "thank", "excite", "kneel", "tease", "fix", "baseball", "taste", "jobless", "power", "axiomatic", "card", "bouncy", "drip", "silky", "quickest", "tendency", "ahead", "electric", "rock", "foolish", "command", "understood", "ossified", "shoe", "clover", "ashamed", "fixed", "chemical", "blush", "harmony", "lively", "high", "home", "pizzas", "giddy", "carpenter", "fumbling", "mixed", "tenuous", "rain", "satisfy", "strap", "shame", "silent", "hammer", "vanish", "harbor", "sign", "long", "mammoth", "detect", "second-hand", "talk", "pancake", "play", "mark", "zippy", "finger", "behave", "pollution", "stitch", "watch", "wrestle", "energetic", "troubled", "tiresome", "purpose", "approve", "consider", "multiply", "temper", "natural", "corn", "sophisticated", "cheer", "callous", "punish", "disagree", "whirl", "garrulous", "moon", "bitter", "amuse", "hole", "glib", "increase", "drag", "theory", "short", "scientific", "rely", "poor", "education", "satisfying", "versed", "silly", "shape", "industry", "delight", "bashful", "bump", "grubby", "roasted", "program", "gifted", "painful", "hammer", "bird", "same", "endurable", "futuristic", "trick", "cake", "choke", "meddle", "test", "sweltering", "anger", "memory", "pin", "fine", "spill", "lettuce", "appreciate", "plausible", "daughter", "parched", "delicate", "well-groomed", "high-pitched", "staking", "border", "pen", "bone", "plate", "heap", "frame", "two", "wandering", "suspect", "four", "jumpy", "spicy", "pray", "point", "heat", "spiritual", "milky", "massive", "mailbox", "bruise", "concern", "stomach", "bright", "robust", "place", "complete", "flame", "hose", "sack", "twist", "skinny", "church", "collect", "abounding", "fetch", "fortunate", "joke", "look", "splendid", "profuse", "decorous", "boot", "texture", "fantastic", "point", "add", "entertaining", "torpid", "meal", "sin", "tidy", "animal", "trick", "yielding", "view", "instinctive", "voice", "book", "strip", "spring", "dashing", "bless", "cemetery", "untidy", "place", "kick", "defective", "jog", "scarecrow", "peaceful", "nut", "unused", "record", "hop", "nest", "awesome", "park", "screw", "kick", "want", "bizarre", "stay", "moor", "labored", "volatile", "seat", "somber", "thread", "untidy", "badge", "extend", "nerve", "classy", "plastic", "suit", "obedient", "rat", "thin", "hang", "explain", "describe", "cycle", "shy", "throat", "toad", "bounce", "fearless", "cross", "twig", "rough", "air", "grey", "return", "ghost", "match", "reduce", "selective", "doubtful", "puncture", "move", "last", "regular", "homely", "demonic", "liquid", "spark", "dress", "excellent", "nervous", "curly", "voiceless", "straw", "zealous", "end", "umbrella", "fair", "repeat", "worried", "reward", "visit", "mundane", "shock", "carry", "glue", "pinch", "plane", "sparkling", "cough", "glow", "stretch", "insect", "freezing", "treat", "plug", "skate", "girls", "prose", "aboriginal", "testy", "maid", "creature", "embarrass", "bewildered", "private", "thirsty", "charge", "flat", "pop", "puny", "avoid", "teeny", "limping", "shrug", "stretch", "bike", "happy", "scold", "harsh", "murder", "ajar", "wander", "territory", "spot", "symptomatic", "scene", "alike", "concerned", "accurate", "pat", "cook", "volcano", "insurance", "toothbrush", "actually", "count", "old-fashioned"}
