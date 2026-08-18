FROM archlinux:latest

RUN pacman -Syyu --needed --noconfirm \
    base-devel sudo git \
    pipewire wireplumber pipewire-pulse pipewire-alsa pipewire-jack \
    mesa libva-mesa-driver vulkan-radeon \
    && pacman -Scc --noconfirm

ARG UID=1000
ARG GID=1000
ARG USER=arch

RUN groupadd -g $GID $USER \
    && useradd -mG wheel -s /bin/bash -u $UID -g $GID $USER

RUN echo "%wheel ALL=(ALL:ALL) NOPASSWD:ALL" >>/etc/sudoers

USER $USER

WORKDIR /home/$USER

RUN git clone https://aur.archlinux.org/yay-bin.git \
    && cd yay-bin \
    && makepkg -si --noconfirm

CMD ["bash"]
