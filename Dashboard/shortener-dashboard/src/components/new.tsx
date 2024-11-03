import { useState } from "react";

const NewLink = (setShortenedUrls: any) => {
  const [newActive, setNewActive] = useState(false);
  const [newUrl, setNewUrl] = useState({ url: "", slug: "" });

  const handleUrlChange = (e: any) => {
    setNewUrl({ ...newUrl, url: e.target.value });
  };

  const handleSlugChange = (e: any) => {
    setNewUrl({ ...newUrl, slug: e.target.value });
  };

  const setNewSlug = async () => {
    const res = await fetch("http://localhost:5000/set", {
      method: "POST",
      body: JSON.stringify({ url: newUrl.url, slug: newUrl.slug }),
    });

    if (res.ok) {
      setNewActive(false);
      setNewUrl({ url: "", slug: "" });
      setShortenedUrls((prevUrls: any) => [...prevUrls, newUrl]);
    }
  };

  return (
    <div className="col-span-full">
      {newActive ? (
        <div className="flex space-x-4">
          <span className="flex flex-col">
            <label className="mb-1 text-gray-700">URL</label>
            <input
              type="text"
              className="border-b-2 border-gray-300 focus:border-blue-500 focus:outline-none transition-colors duration-200"
              value={newUrl.url}
              onChange={handleUrlChange}
            />
          </span>

          <span className="flex flex-col">
            <label className="mb-1 text-gray-700">Slug</label>
            <input
              type="text"
              className="border-b-2 border-gray-300 focus:border-blue-500 focus:outline-none transition-colors duration-200"
              value={newUrl.slug}
              onChange={handleSlugChange}
            />
          </span>

          <button onClick={setNewSlug}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6 text-green-500"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M12 6v6m0 0v6m0-6h6m-6 0H6"
              />
            </svg>
          </button>
        </div>
      ) : (
        <>
          <button
            onClick={() => setNewActive(true)}
            className="p-2 px-4 rounded-md bg-gray-100 border border-gray-300 "
          >
            New
          </button>
        </>
      )}
    </div>
  );
};

export default NewLink;
