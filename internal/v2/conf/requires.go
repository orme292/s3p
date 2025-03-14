package conf

import (
    "strings"

    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func MapValueSafe(key string, confMap map[string]string) string {
    if v, ok := confMap[key]; ok {
        return v
    }
    return EmptyStr
}

func RequireKeys(required []string, confMap map[string]string) (missing []string) {
    count := make(map[string]int)
    for k := range required {
        if _, ok := confMap[required[k]]; ok {
            if confMap[required[k]] != EmptyStr {
                count[confMap[required[k]]]++
            }
        }
    }
    for k := range count {
        missing = append(missing, k)
    }
    return missing
}

func RequireOneOfKey(valid []string, confMap map[string]string) (bool, []string) {
    var result []string
    for _, str := range valid {
        if _, ok := confMap[str]; ok {
            if confMap[str] != EmptyStr {
                result = append(result, str)
            }
        }
    }
    return len(result) == 1, result
}

func RequireOneOfKeyValue(key string, validVals []string, confMap map[string]string) (bool, string) {
    if val := MapValueSafe(key, confMap); val != EmptyStr {
        for _, str := range validVals {
            if val == str {
                return true, val
            }
        }
    }
    return false, EmptyStr
}

type FormatMapFunc func(key, val string) (string, string)

var PlainMap = func(key, val string) (string, string) {
    return lowerString(stripSpace(key)), lowerString(stripSpace(val))
}

func FormatMap(confMap map[string]string, f FormatMapFunc) map[string]string {
    newMap := make(map[string]string)
    for key, val := range confMap {
        newKey, newVal := f(key, val)
        newMap[newKey] = newVal
    }
    return newMap
}

type FormatStringSliceFunc func(val string) string

var PlainStringSlice = func(val string) string {
    return lowerString(stripSpace(val))
}

func FormatStringSlice(strings []string, f FormatStringSliceFunc) []string {
    var newStrings []string
    for _, str := range strings {
        newStrings = append(newStrings, f(str))
    }
    return newStrings
}

func lowerString(s string) string {
    cased := cases.Lower(language.English)
    return cased.String(s)
}

func stripSpace(s string) string {
    return strings.TrimSpace(s)
}
