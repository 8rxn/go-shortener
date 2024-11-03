import React from "react";
import { SetStateAction, useEffect, useState } from "react";
import NewLink from "./new";

const Links = () => {
  const [shortenedUrls, setShortenedUrls] = useState([]);
  const [editingSlug, setEditingSlug] = useState(null);
  const [newUrl, setNewUrl] = useState("");

  const host = "https://sh.aryn.wtf/";

  useEffect(() => {
    fetchAllUrl();
  }, []);

  const fetchAllUrl = async () => {
    try {
      const response = await fetch("http://localhost:5000/get-all", {
        method: "POST",
      });
      if (!response.ok) {
        throw new Error("Failed to fetch URLs");
      }
      const data = await response.json();
      setShortenedUrls(data);
    } catch (error) {
      console.error(error);
    }
  };

  const handleDelete = async (slug: any) => {
    try {
      const response = await fetch(`http://localhost:5000/delete`, {
        method: "POST",
        body: JSON.stringify({ slug }),
      });
      if (response.ok) {
        setShortenedUrls((prevUrls) =>
          prevUrls.filter((url: any) => url.slug !== slug)
        );
      } else {
        throw new Error("Failed to delete URL");
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleEdit = (
    slug: SetStateAction<null>,
    currentUrl: SetStateAction<string>
  ) => {
    setEditingSlug(slug);
    setNewUrl(currentUrl);
  };

  const handleUpdate = async (slug: any) => {
    try {
      const response = await fetch(`http://localhost:5000/edit/${slug}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url: newUrl }),
      });
      if (response.ok) {
        // Update the URL in state
        setShortenedUrls((prevUrls: any) =>
          prevUrls.map((url: any) =>
            url.slug === slug ? { ...url, url: newUrl } : url
          )
        );
        setEditingSlug(null);
        setNewUrl("");
      } else {
        throw new Error("Failed to update URL");
      }
    } catch (error) {
      console.error(error);
      // Optionally, display an error message to the user
    }
  };

  const handleCancelEdit = () => {
    setEditingSlug(null);
    setNewUrl("");
  };

  return (
    <div className="mt-4 grid grid-cols-3 sm:grid-cols-6 border border-gray-500 rounded-md w-fit p-4 gap-3 ">
      {shortenedUrls?.map((url: { slug: string; url: string }, i) => (
        <React.Fragment key={url.slug}>
          <span>{i + 1}</span>
          <a href={`${host}${url.slug}`} className="hover:underline col-span-2">
            {`${host}${url.slug}`}
          </a>
          <a
            href={url.url.startsWith("http") ? url.url : `http://${url.url}`}
            className="hover:underline col-span-2"
          >
            {url.url}
          </a>
          <span className="p-1 flex justify-center gap-2">
            {editingSlug === url.slug ? (
              <>
                <input
                  type="text"
                  value={newUrl}
                  onChange={(e) => setNewUrl(e.target.value)}
                  className="border rounded p-1"
                />
                <button
                  onClick={() => handleUpdate(url.slug)}
                  className="bg-green-500 text-white p-1 rounded-md"
                >
                  Save
                </button>
                <button
                  onClick={handleCancelEdit}
                  className="bg-gray-500 text-white p-1 rounded-md"
                >
                  Cancel
                </button>
              </>
            ) : (
              <>
                <button
                  onClick={() => handleDelete(url.slug)}
                  className="bg-red-500 text-white p-1 rounded-md"
                >
                  Delete
                </button>
              </>
            )}
          </span>
        </React.Fragment>
      ))}

      <NewLink setShortenedUrls={setShortenedUrls}></NewLink>
    </div>
  );
};

export default Links;
