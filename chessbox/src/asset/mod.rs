use std::borrow::Cow;

use bevy::{
    asset::load_internal_binary_asset,
    prelude::{App, AudioSource, HandleUntyped, Image, Plugin},
    reflect::TypeUuid,
    render::texture::{CompressedImageFormats, ImageSampler, ImageType},
    text::Font,
};

fn font_loader(bytes: &[u8], _: Cow<str>) -> Font {
    Font::try_from_bytes(bytes.to_vec()).expect("could not load font")
}

fn mp3_loader(bytes: &[u8], _: Cow<str>) -> AudioSource {
    AudioSource {
        bytes: bytes.into(),
    }
}

fn image_loader(bytes: &[u8], _: Cow<str>) -> Image {
    let mut image = Image::from_buffer(
        bytes,
        ImageType::Extension("Gif"),
        CompressedImageFormats::NONE,
        true,
    )
    .expect("could not load image");

    let mut image_descriptor = ImageSampler::nearest_descriptor();
    image_descriptor.label = Some("pieces_image");
    image.sampler_descriptor = ImageSampler::Descriptor(image_descriptor);
    image
}

pub struct AssetPlugin;

pub const FONT_HANDLE: HandleUntyped =
    HandleUntyped::weak_from_u64(Font::TYPE_UUID, 436509473926038);
pub const PIECE_HANDLE: HandleUntyped =
    HandleUntyped::weak_from_u64(Image::TYPE_UUID, 510291613494514);
pub const CLICK_SOUND_HANDLE: HandleUntyped =
    HandleUntyped::weak_from_u64(AudioSource::TYPE_UUID, 510291613494514);

impl Plugin for AssetPlugin {
    fn build(&self, app: &mut App) {
        load_internal_binary_asset!(app, FONT_HANDLE, "font/wenkai.ttf", font_loader);
        load_internal_binary_asset!(app, PIECE_HANDLE, "image/background.gif", image_loader);
        load_internal_binary_asset!(app, CLICK_SOUND_HANDLE, "sound/click.mp3", mp3_loader);
    }
}
