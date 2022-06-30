import os
from git import Repo

os.system("go mod tidy")

repo = Repo.init(None)
repo.index.add(repo.untracked_files)
repo.index.commit("Init")
