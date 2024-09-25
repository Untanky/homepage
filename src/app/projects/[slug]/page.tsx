import React from 'react';

export default function ProjectPage({ params }: { params: { slug: string } }) {
  return (
    <h1 className="text-4xl font-bold">
      Project Page:
      {params.slug}
    </h1>
  );
};
