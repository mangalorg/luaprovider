---> name: Name
---> description: Description
---> version: 0.1.0
---> website: https://github.com/mangalorg/luaprovider

---@alias Manga { id: string, title: string, url: string?, cover: string?, banner: string?, anilist_search: string?, [any]: any }
---@alias Volume { number: number, [any]: any }
---@alias Chapter { title: string, url: string?, number: number?, [any]: any }
---@alias Page { url: string, headers: table<string, string>?, cookies: table<string, string>?, extension: string?}

local sdk = require('sdk')

---@param query string
---@return Manga[]
function SearchMangas(query)
  return {}
end

---@param manga Manga
---@return Volume[]
function MangaVolumes(manga)
  return {}
end

---@param volume Volume
---@return Chapter[]
function VolumeChapters(volume)
  return {}
end

---@param chapter Chapter
---@return Page[]
function ChapterPages(chapter)
  return {}
end
